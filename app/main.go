package main

import (
	"log"
	"net/http"
	"net/url"
	"reflect"
	"sync"

	"github.com/naiba/com"
	rwn "github.com/naiba/remote-work-news"
	"github.com/naiba/remote-work-news/crawlers"
	"github.com/robfig/cron"
)

func main() {
	var crawlerTargetForgin = []crawlers.Crawler{
		&crawlers.VueJobsCrawler{},
		&crawlers.ZipRecruiterCrawler{},
		&crawlers.StackOverFlowCrawler{},
	}

	var crawlerTargetChina = []crawlers.Crawler{
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuGolang,
		},
		&crawlers.RubyChinaCrawler{},
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuLaravel,
		},
		&crawlers.YizaoyiwanCrawler{},
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuPHP,
		},
		&crawlers.SegmentFaultCrawler{},
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuPython,
		},
		&crawlers.YuanChengDotWorkCrawler{},
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuVueJS,
		},
	}

	// 抓取计划
	c := cron.New()
	c.AddFunc("0 0 0 * * *", func() {
		do(crawlerTargetChina)
	})
	c.AddFunc("0 0 12 * * *", func() {
		do(crawlerTargetForgin)
	})
	c.Start()
}

func do(c []crawlers.Crawler) {
	var errorMsg []byte
	var allNews []rwn.News
	var l sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(c))
	for i := 0; i < len(c); i++ {
		go func(i int) {
			news, err := c[i].FetchNews()
			if err != nil {
				errorMsg = append(errorMsg, ("- " + reflect.TypeOf(c).String() + ":" + err.Error() + "\n")...)
			}
			err = c[i].FillContent(news)
			if err != nil {
				errorMsg = append(errorMsg, ("- " + reflect.TypeOf(c).String() + ":" + err.Error() + "\n")...)
			}
			l.Lock()
			allNews = append(allNews, news...)
			l.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < len(allNews); i++ {
		allNews[i].Hash = com.MD5(allNews[i].URL)
		rwn.DB.Save(allNews[i])
	}
	if len(errorMsg) > 0 {
		serverChan("「远程工作」抓取错误", string(errorMsg))
	}
}

func serverChan(title, content string) {
	log.Println(title, content)
	params := url.Values{
		"text": {title},
		"desp": {content},
	}
	http.PostForm("https://sc.ftqq.com/"+rwn.C.ServerChan+".send", params)
}
