package crawlers

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

const (
	learnKuBase     = "https://learnku.com"
	learnKuSelector = "div.readme"
	// LearnKuLaravel laravel
	LearnKuLaravel = "/laravel/c/jobs?filter=recent"
	// LearnKuGolang golang
	LearnKuGolang = "/golang/c/jobs?filter=recent"
	// LearnKuPHP PHP
	LearnKuPHP = "/php/c/jobs?filter=recent"
	// LearnKuVueJS VueJS
	LearnKuVueJS = "/vuejs/c/jobs?filter=recent"
	// LearnKuPython Python
	LearnKuPython = "/python/c/jobs?filter=recent"
)

// LearnKuCrawler LearnKu
type LearnKuCrawler struct {
	LearnKuChannel string
}

// FetchNews 抓取列表
func (k *LearnKuCrawler) FetchNews() ([]rwn.News, error) {
	doc, err := getDocFromURL(learnKuBase + k.LearnKuChannel)
	if err != nil {
		return nil, err
	}
	var news []rwn.News
	doc.Find("div.simple-topic").Each(func(i int, s *goquery.Selection) {
		var newsItem rwn.News
		newsItem.MediaID = 3
		pusherLink := s.Find("div.user-avatar a").First()
		newsItem.PusherLink = pusherLink.AttrOr("href", "")
		newsItem.Pusher = pusherLink.Children().First().AttrOr("alt", "")
		titleAndLink := s.Find("a.topic-title-wrap").First()
		newsItem.Title = titleAndLink.Find("span.topic-title").First().Text()
		newsItem.URL = titleAndLink.AttrOr("href", "")
		newsItem.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", s.Find("abbr.timeago").First().AttrOr("title", ""))
		if matchRemoteChinese.MatchString(newsItem.Title) {
			news = append(news, newsItem)
		}
	})
	return news, nil
}

// FillContent 抓取内容
func (k *LearnKuCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, learnKuSelector)
}
