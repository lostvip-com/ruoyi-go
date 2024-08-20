package lv_chrome

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"time"
)

// VisitWeb 任务 主要用来设置cookie ，获取登录账号后的页面
func VisitWeb(ctx context.Context, url string, domain string, cookies map[string]string) error {
	//创建一个chrome任务
	err := chromedp.Run(ctx,
		//ActionFunc是一个适配器，允许使用普通函数作为操作。
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 设置Cookie存活时间
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// 添加Cookie到chrome
			if cookies != nil {
				for k, v := range cookies {
					//SetCookie使用给定的cookie数据设置一个cookie； 如果存在，可能会覆盖等效的cookie。
					err := network.SetCookie(k, v).
						WithExpires(&expr).
						// 设置cookie作用的站点
						WithDomain(domain). //访问网站主体
						// 设置httponly,防止XSS攻击
						WithHTTPOnly(true).
						//Do根据提供的上下文执行Network.setCookie。
						Do(ctx)
					if err != nil {
						return err
					}
				}
			}
			return nil
		}),
		// 跳转指定的url地址
		chromedp.Navigate(url),
	)
	return err
}
