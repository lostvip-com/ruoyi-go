package lv_tpl

import (
	"bytes"
	"github.com/lv_framework/utils/lv_secret"
	"os"
	"strings"
	"text/template"
)

// 读取模板
func ParseTemplate(templateName string, data interface{}) (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(cur + "/template/" + templateName)
	if err != nil {
		return "", err
	}
	templateStr := string(b)

	tmpl, err := template.New(templateName).Parse(templateStr) //建立一个模板，内容是"hello, {{OssUrl}}"
	if err != nil {
		return "", nil
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, data) //将string与模板合成，变量name的内容会替换掉{{OssUrl}}
	if err != nil {
		return "", nil
	}
	return buffer.String(), nil
}

// 读取模板
func ParseTemplateStr(templateStr string, data interface{}) (string, error) {
	templateName := lv_secret.Md5(templateStr)
	tmpl, err := template.New(templateName).Parse(templateStr) //建立一个模板，内容是"hello, {{OssUrl}}"
	if err != nil {
		return "", nil
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, data) //替换模板变量
	if err != nil {
		return "", nil
	}
	str := strings.ReplaceAll(buffer.String(), "\n", " ")
	return str, nil
}
