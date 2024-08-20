package entity

import (
	"encoding/json"
	"fmt"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_net"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
)

type INBPoster struct {
	req Req4INB
}

type Req4INB struct {
	KeyWord      string `json:"key_word"`
	IndustryCode string `json:"industry_code"`
	PageSize     int    `json:"page_size"`
	PageNumber   int    `json:"page_number"`
	AreaKey      string `json:"area_key"`
}

// 商圈|无小区数据版本
func (e *INBPoster) Post4ZhengCe(filePath string, req Req4INB, auth string) (string, error) {
	e.req = req
	header := make(map[string]string, 0)
	header["Connection"] = "keep-alive"
	header["Accept"] = "*/*"
	header["Origin"] = "https://www.chanyeos.com"
	//关键头 如果没有 则返回 错误
	header["Referer"] = "https://www.chanyeos.com/smart-ke/"
	//header["Accept-Encoding"] = "gzip, deflate"
	//header["Accept-Language"] = "zh-CN,zh;q=0.9"
	header["Authorization"] = auth

	url := "https://www.chanyeos.com/industry_knowledge_engine/v2/industryChannel/find_policy_list"

	data, _ := json.Marshal(req)
	ret, err := lv_net.PostJsonAndHeader(url, data, header)
	lv_err.HasErrAndPanic(err)
	fmt.Println("================ over =========================" + string(ret))

	//	data := `
	//{
	//	"status": "00000",
	//	"message": "请求成功",
	//	"data": {
	//        "rank": ["1"],
	//		"count": 17215,
	//		"dataSource": [
	//			{
	//				"title": "1在1",
	//
	//				"type": "123"
	//			}
	//		]
	//	}
	//}
	//`
	//	ret := []byte(data)
	var result = make(map[string]any, 0)
	err = json.Unmarshal(ret, &result)
	if err != nil {
		panic(ret)
	}
	err = e.wirteToExcellZhengCe(filePath, result)
	return filePath, err
}

func (e *INBPoster) wirteToExcellZhengCe(filePath string, result map[string]any) error {
	data := result["data"].(map[string]any)
	//fmt.Println(data)
	dataSource := data["dataSource"]
	//fmt.Println("===>dataSource:", dataSource)
	arr := dataSource.([]any)

	listRows := make([][]string, 0)
	for i, item := range arr {
		j := (e.req.PageNumber-1)*20 + i
		mp := item.(map[string]any)
		fmt.Println(j, "=============", mp["title"])
		row := make([]string, 10)
		row[0] = mp["index"].(string)
		row[1] = mp["title"].(string)
		row[2] = mp["location"].(string)
		row[3] = e.getStrByArr(mp["rank"])
		row[4] = mp["type"].(string)
		row[5] = e.getStrByArr(mp["support_way"])
		row[6] = e.getStrByArr(mp["support_object"])
		row[7] = mp["is_declare"].(string)
		row[8] = mp["link"].(string)
		listRows = append(listRows, row)
	}

	var heads []string = nil
	if e.req.PageNumber == 1 {
		heads = []string{"序号", "政策名称", "区域", "政策层级", "政策类型", "支持方式", "支持对象", "申报类型", "连接地址"}
	}
	err := lv_office.Write2Xls(filePath, "政策列表", heads, listRows)
	return err
}

func (e *INBPoster) getStrByArr(a any) string {
	var str = ""
	if a == nil {
		return str
	}
	arr := a.([]any)
	for _, s := range arr {
		str = str + " " + s.(string)
	}
	return str
}
