package crawlers

import (
	rwn "github.com/naiba/remote-work-news"
)

// VueJobsCrawler VueJobs
type VueJobsCrawler struct {
}

type vueJobsData struct {
	Title       string
	Description string
	Route       string
	Company     struct {
		Name  string
		Route string
	}
}

// FetchNews 抓取列表
func (v *VueJobsCrawler) FetchNews() ([]rwn.News, error) {
	_, body, errs := request.Get("https://vuejobs.com/remote-vuejs-jobs").End()
	if errs != nil {
		return nil, errs[0]
	}
	var news []rwn.News

	news = append(news, item)
	return news, nil
}

// FillContent 抓取内容
func (v *VueJobsCrawler) FillContent(news []rwn.News) error {
	return nil
}
