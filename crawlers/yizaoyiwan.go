package crawlers

import (
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

// YizaoyiwanCrawler 一早一晚抓取
type YizaoyiwanCrawler struct {
}

const yzywBase = "https://yizaoyiwan.com"

// FetchNews 抓取列表
func (y *YizaoyiwanCrawler) FetchNews() ([]rwn.News, error) {
	_, body, errs := request.Get(yzywBase + "/categories/employer").End()
	if errs != nil {
		return nil, errs[0]
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var news []rwn.News
	doc.Find("li.discussion.list-group-item>div.media div.media-body").Each(func(i int, s *goquery.Selection) {
		var newsItem rwn.News
		newsItem.MediaID = 1
		titleAndLink := s.Find("div.media-heading a").First()
		newsItem.Title = titleAndLink.Text()
		newsItem.URL = yzywBase + titleAndLink.AttrOr("href", "")
		s.Find("div.media-meta").Children().Each(func(i int, s *goquery.Selection) {
			switch i {
			case 1:
				newsItem.Pusher = s.Text()
				newsItem.PusherLink = yzywBase + s.AttrOr("href", "")
			case 2:
				t, ok := s.Attr("datetime")
				if ok {
					newsItem.CreatedAt, _ = time.Parse("2006-01-02T15:04:05-07:00", t)
				}
			}
		})
		news = append(news, newsItem)
	})
	return news, nil
}

// FillContent 抓取内容
func (y *YizaoyiwanCrawler) FillContent(news []rwn.News) error {
	for i := 0; i < len(news); i++ {
		_, body, errs := request.Get(news[i].URL).End()
		if errs != nil {
			return errs[0]
		}
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
		if err != nil {
			return err
		}
		news[i].Content = doc.Find("div.post-content").First().Text()
		time.Sleep(time.Second * 10)
	}
	return nil
}
