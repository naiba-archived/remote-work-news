package main

import (
	"log"

	"github.com/naiba/remote-work-news/crawlers"
)

func main() {
	var crawlerArray = []crawlers.Crawler{
		&crawlers.RubyChinaCrawler{},
		&crawlers.YizaoyiwanCrawler{},
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
