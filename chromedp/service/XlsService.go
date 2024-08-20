package service

//
//import (
//	util_xls "crawler/util"
//	"fmt"
//	"github.com/spf13/cast"
//	"github.com/xuri/excelize/v2"
//	"strings"
//	"time"
//)
//
//type XlsService struct {
//}
//
///**
// * 先在excell中自增生成参量号，再跟mysql对比，没有的就插入，有的就替换excell中的
// */
//func (this *XlsService) SelectParamsToXls(filePath string, sheetNames []string) {
//	var keys = []string{"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O"}
//	f, _ := excelize.OpenFile(filePath)
//	for _, sheetName := range sheetNames {
//		//names := []string{"监控类型","监控类型ID","参量名称","参量ID","类型名称"}
//		datas, err := util_xls.ReadFile(f, sheetName, keys)
//		if err != nil {
//			panic(err)
//		}
//
//		svc := ParamService{}
//		for i := 1; i < len(datas); i++ { //第一行是标题，忽略
//			row := datas[i]
//			fmt.Println(cast.ToString(i)+". =================row===================", row)
//			if row == nil {
//				continue
//			}
//			typeName := row["B"].(string)                     //监控类型name
//			typeId := row["C"].(string)                       //监控类型ID
//			paramName := strings.TrimSpace(row["D"].(string)) //参量名称
//			if paramName == "" {
//				fmt.Println("excell缺少参量名！！，忽略此行。。。。。。。。。。。。i:" + cast.ToString(i))
//				continue
//			}
//			paramNO := row["E"].(string) //参量ID
//			if paramNO != "" {
//				fmt.Println("表中已存在参量号，忽略此行。。。。。。。。。。。。。paramNO:" + paramNO)
//				continue
//			}
//			paramTypeName := row["G"].(string) //参量类型名
//			aidiaodo := strings.ToLower(row["M"].(string))
//			unit := row["N"].(string) //单位
//			if aidiaodo == "" {
//				aidiaodo = GetParamType(paramTypeName)
//			}
//			/////////////////////未填写参量号的////////////////////////////////////////////////////////
//			fmt.Printf(typeName, typeId, paramName, paramNO, aidiaodo)
//			//先按名字去查有没有合适的直接用
//			paramNOInt := svc.GetParamNoByName(paramName, aidiaodo, typeId)
//			if "4G信号强度" == paramName {
//				fmt.Println("=======================================================")
//			}
//			if paramNOInt == 0 { //无同名、同类 参量，分配一个新的
//				//没有合适的生成一个新的
//				appParam, _ := svc.getMinParamNotUsed()
//				appParam.ParamType = aidiaodo
//				appParam.Unit = unit
//				appParam.MonitorTypeId = cast.ToInt(typeId)
//				appParam.ParamName = paramName
//				appParam.Remark = "程序导入，" + aidiaodo + "  " + typeName + "  " + typeId
//				appParam.UseFlag = 1
//				appParam.UpdateTime = time.Now()
//				global.GormDB.Save(&appParam)        //更新参量号为已使用
//				key := fmt.Sprintf("%s%d", "E", i+1) //参量号
//				err = f.SetCellValue(sheetName, key, appParam.ParamNo)
//			} else { //有 同名、同类 参量，复用旧的
//				appParam, _ := svc.getParamByNo(paramNOInt)
//				if appParam.UseFlag == 0 { //尚未使用
//					appParam.UseFlag = 1
//					appParam.UpdateTime = time.Now()
//					appParam.Remark += "，复用"
//					global.GormDB.Save(&appParam) //更新参量号为已使用
//				}
//				if paramNO == "" { //更新excell中 未分号的
//					key := fmt.Sprintf("%s%d", "E", i+1) //参量号
//					err = f.SetCellValue(sheetName, key, appParam.ParamNo)
//				}
//			}
//			if err != nil {
//				panic(err)
//			}
//
//		} //end for
//	}
//	timeStr := time.Now().Format("2006-01-02_15_04_05")
//	if err := f.SaveAs(filePath + timeStr + ".xlsx"); err != nil {
//		fmt.Println("XXXXXXXXXXXXXXXXXXXXXXX", err)
//	} else {
//		fmt.Println("===========success======================")
//	}
//
//}
