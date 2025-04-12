package lv_file

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/lostvip-com/lv_framework/lv_log"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}
func PathCreateIfNotExist(dir string) error {
	var err error
	if IsPathExist(dir) == false { //不存在
		err = os.MkdirAll(dir, os.ModePerm)
	}
	return err
}

func DelPath(dir string) error {
	return os.RemoveAll(dir)
}

func IsFileExist(addr string) bool {
	_, err := os.Stat(addr) //可能是文件，也可能是目录
	if err != nil {
		return false
	}
	return true
}

// IsPathExist 判断目录是否存在
func IsPathExist(addr string) bool {
	s, err := os.Stat(addr)
	if err != nil {
		lv_log.Error(err)
		return false
	}
	return s.IsDir()
}

func FileCreate(content bytes.Buffer, name string) (string, error) {
	absPath, err := filepath.Abs(name)
	if err != nil {
		return name, err
	}
	dir := filepath.Dir(absPath)
	err = os.MkdirAll(dir, 0755) // 创建目录及所有必要的上级目录，权限为0755
	if err != nil {
		lv_log.Error(err)
		return name, err
	}
	file, err := os.Create(absPath)
	defer file.Close()
	if err != nil {
		lv_log.Error(err)
		return name, err
	}
	_, err = file.WriteString(content.String())
	if err != nil {
		lv_log.Error(err)
		return name, err
	}
	return absPath, err
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {

	return filepath.Walk(h.Root, h.walkCallback)

}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {

	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	log.Printf("h.OldText: %s \n", h.OldText)
	log.Printf("h.NewText: %s \n", h.NewText)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}

func FileMonitoringById(ctx context.Context, filePth string, id string, group string, hookfn func(context.Context, string, string, []byte)) {
	f, err := os.Open(filePth)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	f.Seek(0, 2)
	for {
		if ctx.Err() != nil {
			break
		}
		line, err := rd.ReadBytes('\n')
		// 如果是文件末尾不返回
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			log.Fatalln(err)
		}
		go hookfn(ctx, id, group, line)
	}
}

// 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// GetCurrentPath 获取当前路径，注意不一定是exe所在目录，而是命令执行的目录
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
