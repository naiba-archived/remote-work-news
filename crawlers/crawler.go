package crawlers

import (
	rwn "github.com/naiba/remote-work-news"
	"github.com/parnurzeal/gorequest"
)

const (
	crawlerRetryTime = 3
	crawlerDelayTime = 5
)

// Crawler 爬虫
type Crawler interface {
	FetchNews() ([]rwn.News, error)
	FillContent(news []rwn.News) error
}

var request *gorequest.SuperAgent

func init() {
	request = gorequest.New().Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3").
		Set("Accept-Language", "zh,en;q=0.9,zh-CN;q=0.8").
		Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36")
}
