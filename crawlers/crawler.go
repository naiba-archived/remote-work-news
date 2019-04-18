package crawlers

const (
	crawlerRetryTime = 3
	crawlerDelayTime = 5
)

type Crawler interface {
	FetchNews() ([]News, error)
	FillContent(news []News) error
}
