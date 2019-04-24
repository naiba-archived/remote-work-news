package crawlers

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

// RubyChinaCrawler RubyChina论坛
type RubyChinaCrawler struct {
}

const rubyChinaBase = "https://ruby-china.org"

// FetchNews 抓取列表
func (r *RubyChinaCrawler) FetchNews() ([]rwn.News, error) {
	doc, err := getDocFromURL(rubyChinaBase + "/jobs")
	if err != nil {
		return nil, err
	}
	var news []rwn.News
	doc.Find("div.infos.media-body").Each(func(i int, s *goquery.Selection) {
		var newsItem rwn.News
		newsItem.MediaID = 2
		titleAndLink := s.Find("div.media-heading a").First()
		newsItem.Title = titleAndLink.Text()
		newsItem.URL = rubyChinaBase + titleAndLink.AttrOr("href", "")
		user := s.Find("div.info>a.user-name").First()
		newsItem.Pusher = user.Text()
		newsItem.PusherLink = rubyChinaBase + user.AttrOr("href", "")
		newsItem.PublishedAt, _ = time.Parse("2006-01-02T15:04:05-07:00", s.Find("abbr.timeago").First().AttrOr("title", ""))
		news = append(news, newsItem)
	})
	return news, nil
}

// FillContent 抓取内容
func (r *RubyChinaCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, "div.topic-detail>div.markdown")
}
