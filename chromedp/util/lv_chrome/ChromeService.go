package lv_chrome

import (
	"context"
	"fmt"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"log"
	"net"
	"sync"
	"time"

	"github.com/chromedp/chromedp"
)

type ChromeService struct {
	AllocatorCtx context.Context
	RunCtx       context.Context
	RunCancel    context.CancelFunc
}

func NewChromeService() *ChromeService {
	chromeService := &ChromeService{}
	return chromeService
}

// StartChromedp use func Cancel() to Close
func (c *ChromeService) StartChromedp(userDataDir string) (context.Context, context.CancelFunc) {
	if !lv_file.IsFileExist(userDataDir) {
		lv_file.Mkdir(userDataDir)
	}
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoDefaultBrowserCheck,                        //不检查默认浏览器
		chromedp.Flag("headless", false),                      // 禁用chrome headless（禁用无窗口模式，那就是开启窗口模式）
		chromedp.Flag("blink-settings", "imagesEnabled=true"), //开启图像界面,重点是开启这个
		chromedp.Flag("ignore-certificate-errors", true),      //忽略错误
		chromedp.Flag("disable-web-security", true),           //禁用网络安全标志
		chromedp.Flag("disable-extensions", true),             //开启插件支持
		chromedp.Flag("disable-default-apps", true),
		chromedp.Flag("enable-automation", false),                       // 防止监测webdriver
		chromedp.Flag("disable-blink-features", "AutomationControlled"), //禁用 blink 特征 作者：知识货栈 https://www.bilibili.com/read/cv24371371/ 出处：bilibili
		chromedp.Flag("user-data-dir", userDataDir),
		chromedp.NoFirstRun, //设置网站不是首次运行
		chromedp.WindowSize(1024, 960),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36"), //设置UserAgent
	)

	if c.IsChromePortOpen() {
		c.AllocatorCtx, _ = chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222/")
	} else {
		c.AllocatorCtx, _ = chromedp.NewExecAllocator(context.Background(), opts...)
	}
	return chromedp.NewContext(c.AllocatorCtx, chromedp.WithLogf(log.Printf))
}

// NavigateWithTimeout 获取整个网页内容
func (c *ChromeService) NavigateWithTimeout(url string, d time.Duration) (html string) {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}
	hasOuter := false
	wg.Add(1)
	time.AfterFunc(d, func() {
		lock.Lock()
		if !hasOuter {
			lv_log.Warn("-----RunNavigateAndOutHtmlWithTimeout---等待超时(%v),执行--AfterFunc\n", d)
			chromedp.Run(c.RunCtx, chromedp.OuterHTML("html", &html))
			hasOuter = true
			wg.Done()
		}
		lock.Unlock()
	})
	go func() {
		// 若执行超时未及时返回，则立即执行页面解析
		chromedp.Run(c.RunCtx, chromedp.Navigate(url))
		lock.Lock()
		if !hasOuter {
			fmt.Printf("-----NavigateAndOutHtmlWithTimeout---指定时间(%v)内响应\n", d)
			chromedp.Run(c.RunCtx, chromedp.OuterHTML("html", &html))
			hasOuter = true
			wg.Done()
		}
		lock.Unlock()
	}()
	wg.Wait()
	return
}

// IsChromePortOpen 检查是否有9222端口，来判断是否运行在linux上
func (c *ChromeService) IsChromePortOpen() bool {
	addr := net.JoinHostPort("", "9222")
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
