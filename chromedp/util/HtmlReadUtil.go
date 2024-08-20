package util_xls

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/lostvip-com/lv_framework/lv_log"
	"strings"
)

// GetSpecialData 得到具体的数据
func GetSpecialData(htmlContent string, selector string) (string, error) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		lv_log.Error(err)
		return "", err
	}
	var str string
	dom.Find(selector).Each(func(i int, selection *goquery.Selection) {
		str = selection.Text()
	})
	return str, nil
}
