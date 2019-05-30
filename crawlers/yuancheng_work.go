package crawlers

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

// YuanChengDotWorkCrawler yuancheng.work
type YuanChengDotWorkCrawler struct {
}

// FetchNews 抓取列表
func (y *YuanChengDotWorkCrawler) FetchNews() ([]rwn.News, error) {
	doc, err := getDocFromURL("https://yuancheng.work")
	if err != nil {
		return nil, err
	}
	var news []rwn.News
	doc.Find("div.job-brief").Each(func(i int, s *goquery.Selection) {
		var newsItem rwn.News
		newsItem.MediaID = 6
		title := s.Find("div.job-body").First()
		link := title.Find("a.job-title").First()
		newsItem.Title = "「" + title.Find("b.text-outstand").Text() + "」" + link.Text()
		newsItem.URL = link.AttrOr("href", "")
		newsItem.Pusher = title.Find("small").First().Text()
		newsItem.PublishedAt, _ = time.Parse("2006-01-02 15:04:05",
			s.Find("div.job-date").AttrOr("date", ""))
		news = append(news, newsItem)
	})
	return news, nil
}

// FillContent 抓取内容
func (y *YuanChengDotWorkCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, "div.main.job")
}
