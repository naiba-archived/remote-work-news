package crawlers

import (
	"time"

	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

const (
	learnKuBase    = "https://learnku.com"
	learnKuLaravel = "/laravel/c/php-jobs?filter=recent"
	learnKuGolang  = "/golang/c/php-jobs?filter=recent"
	learnKuPHP     = "/php/c/php-jobs?filter=recent"
	learnKuVueJS   = "/vuejs/c/php-jobs?filter=recent"
	learnKuPython  = "/python/c/php-jobs?filter=recent"
)

func learnKuFetchNews(learnKuChannel string) ([]rwn.News, error) {
	doc, err := getDocFromURL(learnKuBase + learnKuChannel)
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
		newsItem.CreatedAt, _ = time.Parse("2006-01-02T15:04:05-07:00", s.Find("abbr.timeago").First().AttrOr("title", ""))
		news = append(news, newsItem)
	})
	return news, nil
}
