package main

import (
	"github.com/naiba/remote-work-news/crawlers"
)

/*
每天 12 点抓取一次，依据 URL 去重，保存到数据库
遇到错误 serverChan 通报管理员
*/

func main() {
	var crawlerArray = []crawlers.Crawler{
		&crawlers.VueJobsCrawler{},
		&crawlers.ZipRecruiterCrawler{},
		&crawlers.StackOverFlowCrawler{},
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
	for i := 0; i < len(crawlerArray); i++ {
		news, err := crawlerArray[i].FetchNews()
		if err != nil {
			panic(err)
		}
		err = crawlerArray[i].FillContent(news)
		if err != nil {
			panic(err)
		}
	}
}
