package main

import (
	"chromedp/service"
	"fmt"
)

func main() {
	// 获取参数
	//cityId := flag.String("code", "", "")
	// 解析参数
	//flag.Parse()
	// 参数校验
	//cfg := myconf.GetConfigInstance()
	//lv_log.InitLog("logru.log")
	bd := new(service.PkulawService)
	bd.Test("集成电路")
	//bd := new(service.BaiduService)
	//bd.Test("集成电路")

	//var svc = lv_chrome.NewChromeService()
	//ctx, _ := svc.StartChromedp("d:/test")

	fmt.Println("#############全部完成！！！！")
}
