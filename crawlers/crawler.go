package crawlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
	"github.com/parnurzeal/gorequest"
)

const (
	crawlerRetryTime = 3
	crawlerDelayTime = 6
)

// Crawler 爬虫
type Crawler interface {
	FetchNews() ([]rwn.News, error)
	FillContent(news []rwn.News) error
}

var request *gorequest.SuperAgent

func init() {
	request = gorequest.New().Retry(crawlerRetryTime, time.Second*3, http.StatusGatewayTimeout,
		http.StatusBadGateway, http.StatusBadRequest, http.StatusInternalServerError)
	request.Header = http.Header{
		"accept-language": []string{"zh,en;q=0.9,zh-CN;q=0.8"},
		"cache-control":   []string{"no-cache"},
		"user-agent":      []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36"},
	}
	request.SetDoNotClearSuperAgent(true)
}

func getDocFromURL(url string) (*goquery.Document, error) {
	_, body, errs := request.Get(url).End()
	if errs != nil {
		return nil, errs[0]
	}
	return goquery.NewDocumentFromReader(strings.NewReader(body))
}

func innerFillContent(news []rwn.News, selector string) error {
	for i := 0; i < len(news); i++ {
		doc, err := getDocFromURL(news[i].URL)
		if err != nil {
			return err
		}
		news[i].Content = strings.TrimSpace(doc.Find(selector).First().Text())
		//TODO: remove test code
		return nil
		time.Sleep(time.Second * crawlerDelayTime)
	}
	return nil
}
