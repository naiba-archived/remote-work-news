package crawlers

import (
	"encoding/json"
	"errors"
	"regexp"

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

var vueJobsDataRegexp = regexp.MustCompile(`window\.openings\s=\s(.*)\s*<\/script>\s*<\/head>`)

// FetchNews 抓取列表
func (v *VueJobsCrawler) FetchNews() ([]rwn.News, error) {
	_, body, errs := request.Get("https://vuejobs.com/remote-vuejs-jobs").End()
	if errs != nil {
		return nil, errs[0]
	}
	res := vueJobsDataRegexp.FindStringSubmatch(body)
	if len(res) != 2 {
		return nil, errors.New("VueJobsCrawler needs to update:" + body)
	}
	var news []rwn.News
	var vjd []vueJobsData
	err := json.Unmarshal([]byte(res[1]), &vjd)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(vjd); i++ {
		var item rwn.News
		item.MediaID = 7
		item.Title = vjd[i].Title
		item.URL = vjd[i].Route
		item.Content = vjd[i].Description
		item.Pusher = vjd[i].Company.Name
		item.PusherLink = vjd[i].Company.Route
		if matchRemoteEnglish.MatchString((item.Title)) {
			news = append(news, item)
		}
	}
	return news, nil
}

// FillContent 抓取内容
func (v *VueJobsCrawler) FillContent(news []rwn.News) error {
	return nil
}
