package lv_office

import (
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	sheet        *xlsx.Sheet
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
	filename := strconv.FormatInt(curdate, 10) + ".xlsx"

	filePath := curDir + "/static/upload/" + filename

	err = CreateFilePath(filePath)
	if err != nil {
		log.Printf("%s", err.Error())
		return "", err
	}

	// 创建文件
	file := xlsx.NewFile()
	// 添加新工作表
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	// 向工作表中添加新行
	row := sheet.AddRow()

	// 头部写入
	for _, head := range heads {
		row.AddCell().Value = head
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
	filename := strconv.FormatInt(curdate, 10) + ".xlsx"
	filePath := curDir + "/static/upload/" + filename
	err = CreateFilePath(filePath)
	if err != nil {
		log.Printf("%s", err.Error())
		return "", err
	}

	// 创建文件
	file := xlsx.NewFile()
	// 添加新工作表
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		return "", err
	}
	// 向工作表中添加新行
	row := sheet.AddRow()

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

// 下载Excel
func Write2Xls(filePath string, sheetName string, heads []string, listRows [][]string) error {
	var fileXls *xlsx.File
	if lv_file.IsFileExist(filePath) {
		fileXls, err = xlsx.OpenFile(filePath)
		sheet = fileXls.Sheet[sheetName]
	} else {
		err = CreateFilePath(filePath)
		if err != nil {
			lv_log.Error(err.Error())
			return err
		}
		fileXls = xlsx.NewFile()
		sheet, err = fileXls.AddSheet(sheetName)
		if err != nil {
			return err
		}
	}
	// 头部写入
	if heads != nil {
		row := sheet.AddRow()
		for _, head := range heads {
			cell = row.AddCell()
			cell.Value = head
		}
	}
	// 设置单元格样式
	//sheet.SetColWidth(5, 5, 60) // 设置单元格宽度 0-A 1-B 2-C
	if listRows != nil {
		for _, cells := range listRows {
			row := sheet.AddRow()
			for _, cellData := range cells {
				row.AddCell().Value = cellData
			}
		}
	}
	// 在提供的路径中将文件保存到xlsx文件
	err = fileXls.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}

// WriteMap2Xls 追加excell，由于map无序，必须通过数组传title以保证标题次序
// header := []string{"title", "href", "content"}
// title := map[string]string{"title": "标题", "href": "链接", "content": "内容"}
func WriteMap2Xls(filePath string, sheetName string, heads []string, listRows []map[string]string) error {
	if heads == nil || len(heads) <= 0 {
		panic("标题顺序依赖header，不允许为空!!")
	}
	var fileXls *xlsx.File
	if lv_file.IsFileExist(filePath) {
		fileXls, err = xlsx.OpenFile(filePath)
		sheet = fileXls.Sheet[sheetName]
	} else {
		err = CreateFilePath(filePath)
		if err != nil {
			lv_log.Error(err.Error())
			return err
		}
		fileXls = xlsx.NewFile()
		sheet, err = fileXls.AddSheet(sheetName)
		if err != nil {
			return err
		}
	}
	// 头部写入
	row := sheet.AddRow()
	for _, headV := range heads {
		row.AddCell().Value = headV
	}
	// 设置单元格样式
	//sheet.SetColWidth(5, 5, 60) // 设置单元格宽度 0-A 1-B 2-C
	if listRows != nil {
		for _, rowMap := range listRows {
			row := sheet.AddRow()
			for _, header := range heads { //按header的顺序添加
				row.AddCell().Value = rowMap[header]
			}
		}
	}
	// 在提供的路径中将文件保存到xlsx文件
	err = fileXls.Save(filePath)
	if err != nil {
		return err
	}
	return nil
}
