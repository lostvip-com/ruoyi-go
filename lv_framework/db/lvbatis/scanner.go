package lvbatis

import (
	"bufio"
	"regexp"
	"strings"
)

type Scanner struct {
	line    string
	queries map[string]string
	current string
}

type stateFn func(*Scanner) stateFn

func GetTag(line string) string {
	re := regexp.MustCompile("^\\s*--\\s*name:\\s*(\\S+)")
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return ""
	}
	return matches[1]
}

func initialState(s *Scanner) stateFn {
	if tag := GetTag(s.line); len(tag) > 0 {
		s.current = tag
		return queryState
	}
	return initialState
}

/**
 *　递归解析 tag:sql 到map[string]string{tag:sql}
 */
func queryState(s *Scanner) stateFn {
	if tag := GetTag(s.line); len(tag) > 0 { // 第一步 解析出-- name
		s.current = tag
	} else { // 第二行 尝试按第一行的 tag 存map ,直到发现下一次 tag
		//解析name sql 到map
		s.appendQueryLine()
	}
	return queryState
}

func (s *Scanner) appendQueryLine() {
	current := s.queries[s.current]
	line := strings.Trim(s.line, " \t")
	if len(line) == 0 {
		return
	}

	if len(current) > 0 {
		current = current + "\n"
	}

	current = current + line
	s.queries[s.current] = current
}

func (s *Scanner) Run(io *bufio.Scanner) map[string]string {
	s.queries = make(map[string]string)

	for state := initialState; io.Scan(); {
		s.line = io.Text()
		state = state(s)
	}

	return s.queries
}
