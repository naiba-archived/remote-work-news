package crawlers

import (
	rwn "github.com/naiba/remote-work-news"
)

const (
	crawlerRetryTime = 3
	crawlerDelayTime = 5
)

type Crawler interface {
	FetchNews() ([]rwn.News, error)
	FillContent(news []rwn.News) error
}
