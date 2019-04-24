package main

import (
	"github.com/naiba/remote-work-news/crawlers"
	"github.com/robfig/cron"
	"github.com/naiba/remote-work-news"
	"github.com/naiba/com"
)

/*
遇到错误 serverChan 通报管理员
*/

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

	do(crawlerTargetChina)
	do(crawlerTargetForgin)
	select{}

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
	var allNews []rwn.News
	for i := 0; i < len(c); i++ {
		news, err := c[i].FetchNews()
		if err != nil {
			panic(err)
		}
		err = c[i].FillContent(news)
		if err != nil {
			panic(err)
		}
		allNews = append(allNews,news...)
	}
	for i := 0; i < len(allNews); i++ {
		allNews[i].Hash = com.MD5(allNews[i].URL)
		rwn.DB.Save(allNews[i])
	}
}
