package main

import (
	"github.com/naiba/remote-work-news/crawlers"
)

func main() {
	var crawlerArray = []crawlers.Crawler{
		&crawlers.YizaoyiwanCrawler{},
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
