package crawlers

import (
	rwn "github.com/naiba/remote-work-news"
)

// VueJobsCrawler VueJobs
type VueJobsCrawler struct {
}

var vueJobsData struct {
	Data []struct {
		Key   string
		Title string
	}
}

// FetchNews 抓取列表
func (v *VueJobsCrawler) FetchNews() ([]rwn.News, error) {
	var vjd vueJobsData
	_, _, errs := request.Get("https://vuejobs.com/api/positions/search?search=remote&location=&jobs_per_page=20").
		EndStruct(&vjd)
	if errs != nil {
		return nil, errs[0]
	}
	return nil, nil
}

// FillContent 抓取内容
func (v *VueJobsCrawler) FillContent(news []rwn.News) error {
	return innerFillContent(news, "div#overview-items")
}
