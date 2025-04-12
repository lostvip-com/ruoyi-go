package service

import (
	"bytes"
	"common/global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"system/vo"
)

// ////////////////////////////////////////////////////////////////
// 存入本地的sqlite
//
// ////////////////////////////////////////////////////////////////
type CodeGenService struct {
}

type TplInfo struct {
	pathSrc string
	nameSrc string

	pathDist string
	nameDist string
}

func (e *CodeGenService) ListTpl(dir string) []TplInfo {
	list := []TplInfo{}
	filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		//fmt.Println("file:", filepath.Dir(path), filepath.Base(path))
		if strings.HasSuffix(path, ".tpl") {
			tpl := TplInfo{pathSrc: filepath.Dir(path), nameSrc: filepath.Base(path)}
			list = append(list, tpl)
		} else {
			lv_log.Warn("。。。。。。skip! not end with .tpl " + path)
		}
		return nil
	})
	return list
}

func (e *CodeGenService) PreviewCode(tab *vo.GenTableVO) map[string]map[string]string {
	mapAll := make(map[string]map[string]string)
	//内部模板
	listTpl := e.ListTpl(global.DIR_TPL_CODE_GEN)
	funcs := template.FuncMap{
		"contains": Contains,
	}
	for _, tpl := range listTpl {
		t1, err := template.New(tpl.nameSrc).Funcs(funcs).ParseFiles(tpl.pathSrc + "/" + tpl.nameSrc)
		e.replaceTplVar(tab, &tpl)
		lv_err.HasErrAndPanic(err)
		var b1 bytes.Buffer
		err = t1.Execute(&b1, tab)
		groupKey := filepath.Ext(tpl.nameDist) //获取后缀做为分组key 如 java,go,sql
		if mapAll[groupKey] == nil {           //是否存在此组
			mapAll[groupKey] = make(map[string]string)
		}
		mapAll[groupKey][tpl.nameDist] = b1.String()
	}
	return mapAll
}

func (e *CodeGenService) GenCode(tab *vo.GenTableVO, overwrite bool) {
	//内部模板
	srcTpl := e.ListTpl(global.DIR_TPL_CODE_GEN)
	// 创建函数映射
	funcs := template.FuncMap{
		"contains": Contains,
	}
	for _, tpl := range srcTpl {
		t1, err := template.New(tpl.nameSrc).Funcs(funcs).ParseFiles(tpl.pathSrc + "/" + tpl.nameSrc)
		if err != nil {
			lv_log.Error(err)
			continue
		}
		var b1 bytes.Buffer
		err = t1.Execute(&b1, tab)
		//路径替换
		basePath := lv_file.GetCurrentPath()
		if err != nil {
			lv_log.Error(err)
			continue
		}
		e.replaceTplVar(tab, &tpl)
		targetPath := basePath + tpl.pathDist + "/" + tpl.nameDist
		if overwrite {
			targetPath, err = lv_file.FileCreate(b1, targetPath)
		} else if !lv_file.IsFileExist(targetPath) {
			targetPath, err = lv_file.FileCreate(b1, targetPath)
		}
		lv_log.Info("生成文件:", err, targetPath)
	}

}

// 读取模板
func (svc CodeGenService) LoadTemplate(templateName string, data interface{}) (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(cur + "/resources/template/" + templateName)
	if err != nil {
		return "", err
	}
	templateStr := string(b)

	// 创建函数映射
	funcs := template.FuncMap{
		"contains": Contains,
	}

	tmpl, err := template.New(templateName).Funcs(funcs).Parse(templateStr) //建立一个模板，内容是"hello, {{OssUrl}}"
	if err != nil {
		return "", err
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, data) //将string与模板合成，变量name的内容会替换掉{{OssUrl}}
	return buffer.String(), err
}

/**
 * 替换模板变量中的路径变量
 */
func (e *CodeGenService) replaceTplVar(tab *vo.GenTableVO, tpl *TplInfo) {
	//替换路径中的占位符
	tpl.pathDist = strings.ReplaceAll(tpl.pathSrc, global.DIR_TPL_CODE_GEN, global.DIR_DIST_CODE)

	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.ModuleName}}", tab.ModuleName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.PackageName}}", tab.PackageName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.BusinessName}}", tab.BusinessName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.ClassName}}", tab.ClassName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.FuncName}}", tab.FunctionName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.FunctionName}}", tab.FunctionName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.TbName}}", tab.TbName)
	tpl.pathDist = strings.ReplaceAll(tpl.pathDist, "{{.Tmp}}", "tmp")
	_ = lv_file.PathCreateIfNotExist(tpl.pathDist)
	//替换文件名中的占位符
	tpl.nameDist = tpl.nameSrc[0 : len(tpl.nameSrc)-4]
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.ModuleName}}", tab.ModuleName)
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.PackageName}}", tab.PackageName)
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.BusinessName}}", tab.BusinessName)
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.ClassName}}", tab.ClassName)
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.FuncName}}", tab.FunctionName)
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.FunctionName}}", tab.FunctionName)
	tpl.nameDist = strings.ReplaceAll(tpl.nameDist, "{{.TbName}}", tab.TbName)
}
