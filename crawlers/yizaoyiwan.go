package crawlers

import (
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
	return nil, nil
}

// FillContent 抓取内容
func (y *YizaoyiwanCrawler) FillContent(news []rwn.News) error {
	return nil
}
