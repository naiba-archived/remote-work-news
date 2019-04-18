package crawlers

import (
	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
	"github.com/parnurzeal/gorequest"
)

// YizaoyiwanCrawler 一早一晚抓取
type YizaoyiwanCrawler struct {
}

// FetchNews 抓取列表
func (y *YizaoyiwanCrawler) FetchNews() ([]rwn.News, error) {
	request := gorequest.New()
	_, body, errs := request.Get("https://yizaoyiwan.com/categories/employer").End()
	if errs != nil {
		return nil, errs[0]
	}
	goquery.NewDocumentFromReader(buffio.NewReader())
	return nil, nil
}

// FillContent 抓取内容
func (y *YizaoyiwanCrawler) FillContent(news []rwn.News) error {
	return nil
}
