package lv_chrome

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"strings"
)

// GetChildrenLinkAndText 获取子节点下的所有超链接名称和链接，selector是上级节点
func GetChildrenLinkAndText(ctx context.Context, selector string) (arrTxt, arrLink []string, err error) {
	err = chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			var html string
			err = chromedp.OuterHTML(selector, &html, chromedp.BySearch).Do(ctx)
			if err != nil {
				return err
			}
			doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
			if err != nil {
				return err
			}
			// 查找所有的<a>标签
			doc.Find("a").Each(func(i int, s *goquery.Selection) {
				// 获取href属性
				href, exists := s.Attr("href")
				if !exists {
					href = "href not found"
				}
				text := s.Text()
				// 打印结果
				arrTxt = append(arrTxt, text)
				arrLink = append(arrLink, href)
			})
			return err
		}),
	)
	return
}
