package crawlers

import (
	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

// StackOverFlowCrawler stackoverflow
type StackOverFlowCrawler struct {
}

const stackOverFlowBase = "https://stackoverflow.com"

// FetchNews 抓取列表
func (s *StackOverFlowCrawler) FetchNews() ([]rwn.News, error) {
	doc, err := getDocFromURL(stackOverFlowBase + "/jobs/remote-developer-jobs?sort=p")
	if err != nil {
		return nil, err
	}
	var news []rwn.News
	doc.Find("div.listResults div.-item").Each(func(i int, s *goquery.Selection) {
		var newsItem rwn.News
		newsItem.MediaID = 5
		title := s.Find("a.s-link").First()
		newsItem.Title = title.Text()
		newsItem.URL = stackOverFlowBase + title.AttrOr("href", "")
		newsItem.PublishedAt = calcCreateTime(s.Find("span.ps-absolute.fc-black-500").First().Text())
		newsItem.Pusher = s.Find("div.-company").Text()
		news = append(news, newsItem)
	})
	return news, nil
}

// FillContent 抓取内容
func (s *StackOverFlowCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, "div#overview-items")
}
