package crawlers

import (
	"github.com/PuerkitoBio/goquery"
	rwn "github.com/naiba/remote-work-news"
)

const (
	segmentfaultBase = "https://segmentfault.com"
)

// SegmentFaultCrawler LearnKu
type SegmentFaultCrawler struct {
}

// FetchNews 抓取列表
func (f *SegmentFaultCrawler) FetchNews() ([]rwn.News, error) {
	doc, err := getDocFromURL(segmentfaultBase + "/groups?tab=jobs")
	if err != nil {
		return nil, err
	}
	var news []rwn.News
	doc.Find("div.group__discuss-box").Each(func(i int, s *goquery.Selection) {
		var newsItem rwn.News
		newsItem.MediaID = 4
		s.Find("span a").Each(func(i int, s *goquery.Selection) {
			switch i {
			case 0:
				newsItem.Title = s.Text()
				newsItem.URL = segmentfaultBase + s.AttrOr("href", "")
			case 1:
				newsItem.Pusher = s.Text()
				newsItem.PusherLink = segmentfaultBase + s.AttrOr("href", "")
			}
		})
		if matchRemoteChinese.MatchString((newsItem.Title)) {
			news = append(news, newsItem)
		}
	})
	return news, nil
}

// FillContent 抓取内容
func (f *SegmentFaultCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, "div.content.fmt")
}
