package lv_office

//
//import (
//	"fmt"
//	"github.com/xuri/excelize/v2"
//)
//
///**
// * 按list返回 map列表
// */
//func ReadFile(f *excelize.File, sheetName string, keys []string) ([]map[string]interface{}, error) {
//	listArr := make([]map[string]interface{}, 0)
//
//	rows, err := f.GetRows(sheetName)
//	if err != nil {
//		fmt.Println(err)
//		return nil, err
//	}
//	for rowIndex, _ := range rows {
//		num := rowIndex + 1
//		rowMap := make(map[string]interface{}, 0)
//		for _, colName := range keys {
//			cellValue, err := f.GetCellValue(sheetName, fmt.Sprintf("%s%d", colName, num))
//			rowMap[colName] = cellValue
//			if err != nil {
//				fmt.Println(err)
//				return nil, err
//			}
//		}
//		listArr = append(listArr, rowMap)
//	}
//	return listArr, err
//}
//
//func SetCellValue(fileName string, sheetName string, colChar string, rowNum int, value interface{}) {
//	f, err := excelize.OpenFile(fileName)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 生成流写入对象
//	f.SetCellValue(sheetName, fmt.Sprintf("%s%d", colChar, rowNum), value)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	f.Save()
//}
//
//func SetCellValues(fileName string, sheetName string, dataMap map[string]interface{}) {
//	if len(dataMap) == 0 {
//		return
//	}
//	f, err := excelize.OpenFile(fileName)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	for key, v := range dataMap {
//		f.SetCellValue(sheetName, key, v)
//		if err != nil {
//			fmt.Println("XXXXXXXXXXXXXXX", err)
//			continue
//		}
//	}
//
//	f.Save()
//}
