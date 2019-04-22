package crawlers

import (
	"time"

	rwn "github.com/naiba/remote-work-news"
)

// ZipRecruiterCrawler ZipRecruiter
type ZipRecruiterCrawler struct {
}

type zipRecruiterData struct {
	Data []struct {
		Title       string `json:"title,omitempty"`
		Company     string `json:"company,omitempty"`
		PublishedAt string `json:"published_at,omitempty"`
		URL         string `json:"url,omitempty"`
	} `json:"data,omitempty"`
}

// FetchNews 抓取列表
func (v *ZipRecruiterCrawler) FetchNews() ([]rwn.News, error) {
	var zrd zipRecruiterData
	_, _, errs := request.Get("https://vuejobs.com/api/positions/search?search=remote&location=&jobs_per_page=20").
		EndStruct(&zrd)
	if errs != nil {
		return nil, errs[0]
	}
	var news []rwn.News
	for i := 0; i < len(zrd.Data); i++ {
		var item rwn.News
		item.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", zrd.Data[i].PublishedAt)
		item.Title = zrd.Data[i].Title
		item.URL = zrd.Data[i].URL
		item.Pusher = zrd.Data[i].Company
		news = append(news, item)
	}
	return news, nil
}

// FillContent 抓取内容
func (v *ZipRecruiterCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, "article#job_desc")
}
