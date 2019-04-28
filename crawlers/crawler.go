package crawlers

import (
	"net/http"
	"regexp"
	"strconv"
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
var matchRemoteChinese = regexp.MustCompile(`[远遠]程`)
var matchRemoteEnglish = regexp.MustCompile(`(?i)remote`)
var matchSpace = regexp.MustCompile(`\s+`)

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

var timeRegexp = regexp.MustCompile(`(\d+)(\w)\sago`)

func calcCreateTime(raw string) time.Time {
	item := timeRegexp.FindStringSubmatch(raw)
	origin := time.Now()
	if len(item) == 3 {
		num, _ := strconv.Atoi(item[1])
		switch item[2] {
		case "h":
			origin = origin.Add(time.Hour * time.Duration(-1*num))
		case "d":
			origin = origin.Add(time.Hour * time.Duration(-1*num*24))
		case "m":
			origin = origin.Add(time.Hour * time.Duration(-1*num*24*30))
		case "y":
			origin = origin.Add(time.Hour * time.Duration(-1*num*24*30*12))
		}
	}
	return origin
}

func innerFillContent(news []rwn.News, selector string) error {
	for i := 0; i < len(news); i++ {
		doc, err := getDocFromURL(news[i].URL)
		if err != nil {
			return err
		}
		news[i].Content = doc.Find(selector).First().Text()
		time.Sleep(time.Second * crawlerDelayTime)
	}
	return nil
}

// ClearSpace 清理空格 换行等
func ClearSpace(news []rwn.News) {
	for i := 0; i < len(news); i++ {
		news[i].Title = strings.TrimSpace(news[i].Title)
		news[i].Content = strings.TrimSpace(news[i].Content)
		news[i].Pusher = strings.TrimSpace(news[i].Pusher)
		news[i].Title = matchSpace.ReplaceAllString(news[i].Title, " ")
		news[i].Content = matchSpace.ReplaceAllString(news[i].Content, " ")
	}
}
