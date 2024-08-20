package service

import (
	"chromedp/util/lv_chrome"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/lostvip-com/lv_framework/lv_global"
	"github.com/lostvip-com/lv_framework/lv_log"
	"github.com/lostvip-com/lv_framework/utils/lv_err"
	"github.com/lostvip-com/lv_framework/utils/lv_file"
	"github.com/lostvip-com/lv_framework/utils/lv_office"
	"github.com/lostvip-com/lv_framework/utils/lv_time"
	"github.com/spf13/cast"
	"log"
	"math"
	"math/rand"
	"strings"
	"sync"
	"time"
)

const userDataDir = "d:/chrome-user-data-dir"
const name = "liuguanghan"
const password = "20240725qinghua"
const baseUrl_qinghua = "https://eproxy.lib.tsinghua.edu.cn"

const url_qinghua = "https://eproxy.lib.tsinghua.edu.cn/https/1SjapK4g9QWMKZdqGX5SBIbCowOHrVOPvRyk8/"

//const url_qinghua = "https://www.pkulaw.com/law?channel=SEM-topad"

type PkulawService struct {
	wg sync.WaitGroup
}

func (s PkulawService) GoToSearchIndexPage(ctx context.Context, searchText string) error {
	//进入清华跳转页
	err := chromedp.Run(ctx,
		chromedp.Navigate(url_qinghua),
	)
	//err = chromedp.Run(ctx,
	//	chromedp.WaitVisible(`//*[@id="txtSearch"]`, chromedp.BySearch),
	//	chromedp.SendKeys(`//*[@id="txtSearch"]`, searchText, chromedp.BySearch),
	//)

	for true {
		lv_chrome.Sleep(2000)
		v := lv_chrome.GetValue(ctx, `//*[@id="txtSearch"]`)
		if v == strings.Trim(searchText, " ") {
			break //正常输入了
		}
		lv_log.Warn("等待输入查询条件.....")
	}

	//全文检索
	lv_chrome.Click(ctx, ".areajSelect")
	lv_chrome.Sleep(1000)
	lv_chrome.Click(ctx, `a[areacode="Fulltext"]`)
	lv_chrome.Click(ctx, `btnSearch`, chromedp.ByID)
	lv_chrome.Sleep(2000)
	return err
}

func (s PkulawService) Test(searchTxt string) (string, error) {
	filePath := "./pkulaw_" + lv_time.GetCurrTimeStr("20060102_150405") + ".xlsx"
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
	allocCtx, cancell := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancell()

	ctx, cancell := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancell()
	// 创建超时上下文
	//ctx, cancel := context.WithTimeout(ctx, 1*time.Minute)
	//defer cancel()
	//打开首页
	err := s.GoToSearchIndexPage(ctx, searchTxt)
	//不分组
	err = lv_chrome.Click(ctx, `#rightContent > div.grid-right > div:nth-child(1) > div.list-tool-wrap > div.sort-wrap > div.group > div > label:nth-child(3) > input[type=radio]`)
	if err != nil {
		fmt.Println("不分组 XXXXXXXXXXXXXXXXXXXx", err)
	}
	header := []string{"page", "num", "typeTxt", "title", "href", "content"}
	title := map[string]string{"page": "页码", "num": "序号", "typeTxt": "法规类型", "title": "标题", "href": "链接", "content": "内容"}
	lv_office.WriteMap2Xls(filePath, searchTxt, header, []map[string]string{title})
	//
	typeSel := `#rightContent > div.grid-right > div:nth-child(1) > div.accompanying-wrap > div`

	allTypes, err := lv_chrome.GetNodes(ctx, typeSel)
	lv_err.HasErrAndPanic(err)
	for i, _ := range allTypes { // node中内容是空的
		//点击页面时，元素可能发生成变化，但数量未变，重新获取一次
		if i >= 2 { //只加载前两个,中央法规和地方法规
			break
		}
		typeTxt := "中央法规"
		if i == 1 {
			typeTxt = "地方法规"
		}
		allTypes, err = lv_chrome.GetNodes(ctx, typeSel)
		fmt.Println(cast.ToString(i), "===>", allTypes[i])
		err := lv_chrome.ClickNode(ctx, allTypes[i])
		if err != nil {
			fmt.Println(err)
		}
		//获取节点内容 #subLibraries > li:nth-child(2),css中索引从1开始
		s.SavePagesByType(searchTxt, ctx, typeTxt, filePath, header)
	}
	//分页查询

	lv_chrome.CloseChrome(ctx)
	fmt.Println("all ########### over !!!!!! ")
	return filePath, err
}

func (s PkulawService) SavePagesByType(searchTxt string, ctx context.Context, typeTxt string, filePath string, header []string) {
	// 查数据库时的分页信息
	//txt := lv_chrome.GetText(ctx, `#rightContent > div.grid-right > div:nth-child(1) > div:nth-child(3) > ul > li:nth-child(1) > label`)
	//arrTxt := strings.Split(txt, ` `)
	//pagesTxt := arrTxt[1]
	//arr := strings.Split(pagesTxt, `/`)
	//pageTotal := arr[1]
	startPage := lv_global.Config().GetValueStr("chrome.start-page")
	endPage := lv_global.Config().GetValueStr("chrome.end-page")
	start := 1
	end := 1
	if startPage != "" {
		start = cast.ToInt(startPage)
	}
	if endPage != "" {
		end = cast.ToInt(endPage)
	}
	// 全文件索引的分页信息#rightContent > div.grid-right > div:nth-child(1) > div.list-tool-wrap > div.filtrater-box > span > strong
	pageSize := 20
	totalTxt := lv_chrome.GetText(ctx, `#rightContent > div.grid-right > div:nth-child(1) > div.list-tool-wrap > div.filtrater-box > span > strong`)
	totalRecords := cast.ToInt(totalTxt)
	pagesTotal := int(math.Ceil(float64(totalRecords) / 20.0)) //一共多少页
	if end > pagesTotal {                                      //
		end = pagesTotal
	}
	for pageNum := start; pageNum <= pagesTotal; pageNum++ {
		if pageNum == pagesTotal { //最后一页，记录可能不是20
			pageSize = totalRecords - (pageNum-1)*pageSize
		}

		pageRows, err := s.SaveRowsByPage(ctx, pageNum, pageSize)
		fillDataPageRow(&pageRows, "typeTxt", typeTxt)
		lv_err.HasErrAndPanic(err)
		err = lv_office.WriteMap2Xls(filePath, searchTxt, header, pageRows)
		lv_err.HasErrAndPanic(err)
		//翻页
		if pageNum < pagesTotal { //非最后一页
			ok := toPageOk(ctx, pageNum+1) //下一页
			if !ok {
				break
			}
		}
	}
}

func fillDataPageRow(arr *[]map[string]string, s string, txt string) {
	for _, mp := range *arr {
		mp[s] = txt
	}
}

func toPageOk(ctx context.Context, page int) bool {
	lv_chrome.ScrollToBottom(ctx)
	// 填写页码
	lv_chrome.SendKeys(ctx, `input[name="jumpToNum"]`, cast.ToString(page)) //
	//跳转
	lv_chrome.Sleep(500)
	lv_chrome.Click(ctx, ".jumpBtn") //
	lv_chrome.ScrollToXY(ctx, 0, 500)
	lv_chrome.Sleep(2000)
	return true
}

// SaveRowsByPage 完成一个section后，由于网页限制回退，只能回到搜索页 重新点击section
func (s PkulawService) SaveRowsByPage(ctx context.Context, pageCurr, pageSize int) ([]map[string]string, error) {
	// 导航到目标网页
	mpArr := make([]map[string]string, 0)
	lv_chrome.Sleep(1000)
	var err error
	for row := 1; row <= pageSize; row++ {
		//随机停顿
		randomInt := rand.Intn(5000)
		lv_chrome.Sleep(randomInt)

		mp := make(map[string]string)
		//sel0 := `#rightContent > div.grid-right > div:nth-child(1) > div:nth-child(3) > div:nth-child(` + cast.ToString(row) + `) > div.col > div.t > h4 > a`
		sel0 := `#rightContent > div.grid-right > div:nth-child(1) > div.accompanying-wrap > div:nth-child(` + cast.ToString(row) + `) > div.col > div.t > h4 > a`
		fmt.Println("读取超链接sel : " + sel0)

		mp["title"] = lv_chrome.GetText(ctx, sel0)
		err = chromedp.Run(ctx,
			chromedp.ScrollIntoView(sel0),
			chromedp.Sleep(time.Duration(500)*time.Millisecond),
		)
		node := lv_chrome.GetNode0(ctx, sel0)
		if node != nil {
			href := baseUrl_qinghua + lv_chrome.GetHrefByNode(node)
			mp["href"] = href
			fmt.Println("page: "+cast.ToString(pageCurr)+"=====>", mp)
			mpArr = append(mpArr, mp)
			fmt.Println("end page: " + cast.ToString(pageCurr))
			content, _ := getHrefText(ctx, href)
			mp["content"] = content
		}

	}
	fmt.Println("################################end Page:" + cast.ToString(pageCurr))
	return mpArr, err
}

func getHrefText(ctxParent context.Context, href string) (string, error) {
	var data string
	ctxNew, cancel, err := lv_chrome.NewWindow(ctxParent, href)
	defer cancel()
	data = lv_chrome.GetText(ctxNew, "#divFullText")
	lv_chrome.CloseChrome(ctxNew)
	return data, err
}

//
//func (s PkulawService) GoToSearchIndexPage(ctx context.Context, searchText string) error {
//	//err := chromedp.Run(ctx, chromedp.Navigate("https://www.pkulaw.com/law?channel=SEM-topad"))
//	//err = chromedp.Run(ctx, chromedp.WaitVisible(`//*[@id="txtSearch"]`, chromedp.BySearch))
//	//err = chromedp.Run(ctx, chromedp.SendKeys(`//*[@id="txtSearch"]`, searchText, chromedp.BySearch))
//	//err = chromedp.Run(ctx, chromedp.Click(`//*[@id="btnSearch"]`, chromedp.BySearch))
//	//url := "https://ecollection.lib.tsinghua.edu.cn/databasenav/entrance/detail?mmsid=991021498919603966"
//	//url := "https://www.pkulaw.com/law?channel=SEM-topad"
//
//	//cookie := map[string]string{"test": "test1111111111111111111111111111111"}
//	//lv_chrome.VisitWeb(ctx, url, "pkulaw.com", cookie)
//
//	//err := chromedp.Run(ctx,
//	//	chromedp.Navigate(url),
//	//	chromedp.Click(`#off-campus > a:nth-child(4)`),
//	//	chromedp.Sleep(time.Second*2),
//	//)
//	//ctx, cancel := chromedp.NewContext(ctx)
//	//defer cancel()
//	//lv_chrome.Click(ctx, `/html/body/div/div[2]/div/div/div[6]/div`)
//	//lv_chrome.Sleep(1000)
//	//err = chromedp.Run(ctx,
//	//	//chromedp.Click(`/html/body/div/div[2]/div/div/div[6]/div`),
//	//	chromedp.Click(`#root > div.child___2NTdD > div > div > div.loginBtn___3UI8L > div`),
//	//)
//	//lv_chrome.Sleep(1000)
//	//err = chromedp.Run(ctx,
//	//	chromedp.Click(`#root > div.child___2NTdD > div > div > div.loginBtn___3UI8L > div`, chromedp.ByQuery),
//	//	chromedp.WaitVisible(`#i_user`, chromedp.BySearch),
//	//	chromedp.SendKeys(`#i_user`, "liuguanghan"),
//	//	chromedp.SendKeys(`#i_pass"`, "20240725qinghua"),
//	//	chromedp.Click(`#theform > div:nth-child(6) > a`),
//	//)
//	//lv_chrome.Sleep(1000)
//	//err = chromedp.Run(ctx,
//	//	chromedp.SendKeys(`#i_user`, "liuguanghan"),
//	//	chromedp.SendKeys(`#id="i_pass"`, "20240725qinghua"),
//	//	chromedp.Click(`#theform > div:nth-child(6) > a`),
//	//)
//	//进入清华跳转页
//	err := chromedp.Run(ctx,
//		chromedp.Navigate(url_qinghua),
//	)
//	err = chromedp.Run(ctx,
//		chromedp.WaitVisible(`//*[@id="txtSearch"]`, chromedp.BySearch),
//	)
//	err = chromedp.Run(ctx,
//		chromedp.SendKeys(`//*[@id="txtSearch"]`, searchText, chromedp.BySearch),
//	)
//	lv_chrome.Click(ctx, `btnSearch`, chromedp.ByID)
//	lv_chrome.Sleep(2000)
//	return err
//}
