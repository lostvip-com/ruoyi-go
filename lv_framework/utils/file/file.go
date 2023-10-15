package file

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	// Separator for file system.
	Separator = string(filepath.Separator)
)

var (
	// defaultTrimChars are the characters which are stripped by Trim* functions in default.
	defaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

// CopyFile copies the contents of the file named <src> to the file named
// by <dst>. The file will be created if it does not exist. If the
// destination file exists, all it's contents will be replaced by the contents
// of the source file. The file mode will be copied from the source and
// the copied data is synced/flushed to stable storage.
// Thanks: https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04
func CopyFile(src, dst string) (err error) {
	if src == "" {
		return errors.New("source file cannot be empty")
	}
	if dst == "" {
		return errors.New("destination file cannot be empty")
	}
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer func() {
		if e := in.Close(); e != nil {
			err = e
		}
	}()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	_, err = io.Copy(out, in)
	if err != nil {
		return
	}
	err = out.Sync()
	if err != nil {
		return
	}
	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}
	return
}

// Mkdir creates directories recursively with given <path>.
// The parameter <path> is suggested to be absolute path.
func Mkdir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// in the final element of path; it is empty if there is
// no dot.
//
// Note: the result contains symbol '.'.
func Ext(path string) string {
	ext := filepath.Ext(path)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// Trim strips whitespace (or other characters) from the beginning and end of a string.
// The optional parameter <characterMask> specifies the additional stripped characters.
func Trim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.Trim(str, defaultTrimChars)
	} else {
		return strings.Trim(str, defaultTrimChars+characterMask[0])
	}
}

// SplitAndTrim splits string <str> by a string <delimiter> to an array,
// and calls Trim to every element of this array. It ignores the elements
// which are empty after Trim.
func SplitAndTrim(str, delimiter string, characterMask ...string) []string {
	array := make([]string, 0)
	for _, v := range strings.Split(str, delimiter) {
		v = Trim(v, characterMask...)
		if v != "" {
			array = append(array, v)
		}
	}
	return array
}

// Join joins string array paths with file separator of current system.
func Join(paths ...string) string {
	return strings.Join(paths, Separator)
}

// Exists checks whether given <path> exist.
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// IsFile checks whether given <path> a file, which means it's not a directory.
func IsFile(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// Get returns the value of the environment variable named by the <key>.
// It returns given <def> if the variable does not exist in the environment.
func Get(key string, def ...string) string {
	v, ok := os.LookupEnv(key)
	if !ok && len(def) > 0 {
		return def[0]
	}
	return v
}

// SearchBinary searches the binary <file> in current working folder and PATH environment.
func SearchBinary(file string) string {
	// Check if it's absolute path of exists at current working directory.
	if Exists(file) {
		return file
	}
	return SearchBinaryPath(file)
}

// SearchBinaryPath searches the binary <file> in PATH environment.
func SearchBinaryPath(file string) string {
	array := ([]string)(nil)
	switch runtime.GOOS {
	case "windows":
		envPath := Get("PATH", Get("Path"))
		if strings.Contains(envPath, ";") {
			array = SplitAndTrim(envPath, ";")
		} else if strings.Contains(envPath, ":") {
			array = SplitAndTrim(envPath, ":")
		}
		if Ext(file) != ".exe" {
			file += ".exe"
		}
	default:
		array = SplitAndTrim(Get("PATH"), ":")
	}
	if len(array) > 0 {
		path := ""
		for _, v := range array {
			path = v + Separator + file
			if Exists(path) && IsFile(path) {
				return path
			}
		}
	}
	return ""
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
func Dir(path string) string {
	return filepath.Dir(path)
}

// Create creates file with given <path> recursively.
// The parameter <path> is suggested to be absolute path.
func Create(path string) (*os.File, error) {
	dir := Dir(path)
	if !Exists(dir) {
		Mkdir(dir)
	}
	return os.Create(path)
}
