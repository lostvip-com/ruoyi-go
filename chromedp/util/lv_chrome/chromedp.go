package lv_chrome

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/kb"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/spf13/cast"
	"os"
	"time"
)

func ScrollToTop(ctx context.Context) error {
	lv_log.Info("---------滚动到底部-----------------------------------------------")
	// 定义要执行的actions
	var err error
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools("window.scrollTo(0, 0);", nil),
	)
	return err
}
func ScrollToBottom(ctx context.Context) error {
	lv_log.Info("---------滚动到底部-----------------------------------------------")
	// 定义要执行的actions
	var err error
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools("window.scrollTo(0, document.body.scrollHeight);", nil),
	)
	return err
}

func ScrollToXY(ctx context.Context, x, y int) error {
	lv_log.Info("---------滚动到底部-----------------------------------------------")
	var err error
	err = chromedp.Run(ctx,
		chromedp.EvaluateAsDevTools("window.scrollTo("+cast.ToString(x)+","+cast.ToString(y)+");", nil),
	)
	return err
}

// GetNode0 获取一个节点
func GetNode0(ctx context.Context, selector string) *cdp.Node {
	nodes, err := GetNodes(ctx, selector)
	if err == nil && len(nodes) > 0 {
		return nodes[0]
	}
	return nil
}
func GetNodeValue(node *cdp.Node) string {
	v := node.NodeValue
	return v
}
func GetHrefByNode(node *cdp.Node) string {
	href, _ := node.Attribute("href")
	return href
}

func GetNodes(ctx context.Context, cssClass string) ([]*cdp.Node, error) {
	var cdpNodes []*cdp.Node
	err := chromedp.Run(ctx, chromedp.Nodes(cssClass, &cdpNodes, chromedp.BySearch))
	if err != nil {
		fmt.Println("XXXXXXXXXXXX", err)
	}
	return cdpNodes, err
}

// GetText class必须以点打头
func GetText(ctx context.Context, selector string) string {
	var linkTitle string
	err := chromedp.Run(ctx,
		//chromedp.ScrollIntoView(selector),
		//chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Text(selector, &linkTitle, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println(err)
	}
	return linkTitle
}

// GetValue 获取value值
func GetValue(ctx context.Context, selector string) string {
	var value string
	err := chromedp.Run(ctx,
		chromedp.Value(selector, &value),
	)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

// GetTextByNode class必须以点打头
func GetTextByNode(ctx context.Context, selector string) string {
	var linkTitle string
	err := chromedp.Run(ctx,
		chromedp.ScrollIntoView(selector),
		//chromedp.WaitVisible(selector, chromedp.ByQuery),
		chromedp.Text(selector, &linkTitle, chromedp.ByQuery),
	)
	if err != nil {
		fmt.Println(err)
	}
	return linkTitle
}

func GetAttributeHref(ctx context.Context, cssClass string) string {
	cdpNodes, err := GetNodes(ctx, cssClass)
	if err != nil {
		return ""
	}
	href := cdpNodes[0].AttributeValue("href")
	return href

}

// err = chromedp.Run(ctx,
// chromedp.ScrollIntoView(radio),
// chromedp.WaitVisible(radio),
// chromedp.Click(radio),
// )

// Click  如：css `input[logfunc="分组-不分组"]`
func Click(ctx context.Context, sel string, opts ...chromedp.QueryOption) error {
	err := chromedp.Run(ctx,
		chromedp.ScrollIntoView(sel),
		chromedp.Click(sel, opts...))
	return err
}

func ClickNode(ctx context.Context, node *cdp.Node, opts ...chromedp.MouseOption) error {
	err := chromedp.Run(ctx,
		chromedp.MouseClickNode(node, opts...),
	)
	return err
}

func ClickById(ctx context.Context, id string) error {
	err := chromedp.Run(ctx,
		//chromedp.ScrollIntoView(id),
		chromedp.Click(id, chromedp.ByID))
	return err
}

func SendKeys(ctx context.Context, sel interface{}, v string) error {
	err := chromedp.Run(ctx,
		chromedp.ScrollIntoView(sel),
		chromedp.WaitVisible(sel),
		chromedp.SendKeys(sel, v),
		chromedp.SendKeys(sel, kb.Enter), //回车
	)
	return err
}

func Sleep(millisecond int) {
	time.Sleep(time.Millisecond * time.Duration(millisecond))
}

func CloseChrome(ctx context.Context) {
	// 关闭浏览器窗口
	err := chromedp.Run(ctx)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func NewWindow(ctx context.Context, url string) (context.Context, context.CancelFunc, error) {
	ctx, cancel := chromedp.NewContext(ctx)
	err := chromedp.Run(ctx, chromedp.Navigate(url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
	return ctx, cancel, err
}

//
//func ListAllHref(ctx context.Context, selector string) ([]string, error) {
//	// 定义一个变量来存储所有的超链接
//	var links []string
//	// 获取所有超链接并存储到变量中
//	err := chromedp.Run(ctx,
//		chromedp.Evaluate(`
//            () => {
//                const links = Array.from(document.querySelectorAll('a[href]'));
//                return links;
//            }
//        `, &links), // 执行JavaScript获取所有超链接
//	)
//	return links, err
//}
