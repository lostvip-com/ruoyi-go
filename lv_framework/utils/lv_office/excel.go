package lv_office

import (
	"github.com/tealeg/xlsx"
	"log"
	"lostvip.com/logme"
	"os"
	"path/filepath"
	"robvi/app/common/global"
	"strconv"
	"time"
)

var (
	file         *xlsx.File
	sheet        *xlsx.Sheet
	row          *xlsx.Row
	cell         *xlsx.Cell
	downloadData []interface{}
	err          error
)

// 创建路径
func CreateFilePath(filePath string) error {
	// 路径不存在创建路径
	path, _ := filepath.Split(filePath) // 获取路径
	_, err := os.Stat(path)             // 检查路径状态，不存在创建
	if err != nil || os.IsExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
	}
	return err
}

// 下载Excel
func DownlaodExcel(heads []string, data [][]string) (string, error) {
	// 创建路径
	curDir, err := os.Getwd()

	if err != nil {
		return "", err
	}
	curdate := time.Now().UnixNano()
	filename := strconv.FormatInt(curdate, 10) + ".xls"

	filePath := curDir + "/static/upload/" + filename

	err = CreateFilePath(filePath)
	if err != nil {
		log.Printf("%s", err.Error())
		return "", err
	}

	// 创建文件
	file = xlsx.NewFile()
	// 添加新工作表
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	// 向工作表中添加新行
	row = sheet.AddRow()

	// 头部写入
	for _, head := range heads {
		cell = row.AddCell()
		cell.Value = head
	}
	// 设置单元格样式
	//sheet.SetColWidth(5, 5, 60) // 设置单元格宽度 0-A 1-B 2-C

	// 主体写入数据
	for _, r := range data {
		row = sheet.AddRow()
		for _, c := range r {
			row.AddCell().Value = c
		}
	}
	// 在提供的路径中将文件保存到xlsx文件
	err = file.Save(filePath)
	if err != nil {
		return "", err
	}
	return filename, nil
}

// 下载Excel
func DownlaodExcelByListMap(heads, cols *[]string, listMap *[]map[string]string) (string, error) {
	// 创建路径
	curDir, err := os.Getwd()

	if err != nil {
		return "", err
	}
	curdate := time.Now().UnixNano()
	filename := strconv.FormatInt(curdate, 10) + ".xls"
	filePath := curDir + "/static/upload/" + filename
	if global.GetConfigInstance().IsDebug() {
		logme.Debug("filePath: " + filePath)
	}
	err = CreateFilePath(filePath)
	if err != nil {
		log.Printf("%s", err.Error())
		return "", err
	}

	// 创建文件
	file = xlsx.NewFile()
	// 添加新工作表
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	// 向工作表中添加新行
	row = sheet.AddRow()

	// 头部写入
	for _, head := range *heads {
		cell = row.AddCell()
		cell.Value = head
	}
	// 设置单元格样式
	//sheet.SetColWidth(5, 5, 60) // 设置单元格宽度 0-A 1-B 2-C

	// 主体写入数据
	for _, rowMap := range *listMap {
		row = sheet.AddRow()
		for _, colName := range *cols {
			v := rowMap[colName]
			row.AddCell().Value = v
		}
	}
	// 在提供的路径中将文件保存到xlsx文件
	err = file.Save(filePath)
	if err != nil {
		return "", err
	}
	return filename, nil
}
