package service

import (
	"chromedp/util/lv_chrome"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

type BaiduService struct {
}

//
//chromedp.BySearch // 如果不写，默认会使用这个选择器，类似devtools ctrl+f 搜索，效果等同于`document.querySelector(...)` 去测下
//chromedp.ByID // 只id来选择元素
//chromedp.ByQuery // 根据document.querySelector的规则选择元素，返回单个节点
//chromedp.ByQueryAll // 根据document.querySelectorAll返回所有匹配的节点
//chromedp.ByNodeIP // 检索特定节点(必须先有分配的节点IP)，这个暂时没用过也没看到过例子，如果有例子可以发给我看下
//chromedp.ByJSPath

func (s BaiduService) Test(txt string) {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("enable-automation", false),
	)
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	_ = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate("https://www.baidu.com/"),
		chromedp.SendKeys("//*[@id=\"kw\"]", "落花流水春去也", chromedp.BySearch),
		chromedp.Click(`//*[@id="su"]`, chromedp.BySearch),
	})
	xpathOrSelector := `/html/body`
	//xpathOrSelector := `body`
	arrTxt, links, err := lv_chrome.GetChildrenLinkAndText(ctx, xpathOrSelector)
	fmt.Println("arrTxt==========>", arrTxt, err)
	fmt.Println("links ==========>", links, err)
}
