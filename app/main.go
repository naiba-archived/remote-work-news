package main

import (
	"log"

	"github.com/naiba/remote-work-news/crawlers"
)

func main() {
	var crawlerArray = []crawlers.Crawler{
		&crawlers.YizaoyiwanCrawler{},
	}
	log.Println(crawlerArray[0].FetchNews())
}
