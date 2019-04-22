package main

import (
	"log"

	"github.com/naiba/remote-work-news/crawlers"
)

func main() {
	var crawlerArray = []crawlers.Crawler{
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
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuPython,
		},
		&crawlers.LearnKuCrawler{
			LearnKuChannel: crawlers.LearnKuVueJS,
		},
	}
	for i := 0; i < len(crawlerArray); i++ {
		news, err := crawlerArray[i].FetchNews()
		if err != nil {
			panic(err)
		}
		log.Println(news[0].URL)
		err = crawlerArray[i].FillContent(news)
		if err != nil {
			panic(err)
		}
		log.Println(news[0].Content[:30])
	}
}
