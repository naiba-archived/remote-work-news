package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naiba/com"
	rwn "github.com/naiba/remote-work-news"
	"github.com/naiba/remote-work-news/crawlers"
	"github.com/robfig/cron"
)

var crawling bool

func main() {

	//test code
	// x := &crawlers.VueJobsCrawler{}
	// log.Println(x.FetchNews())
	// os.Exit(0)

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
	_, offset := time.Now().Zone()
	offset /= 3600
	chinaOffset := 12 + offset - 8
	if chinaOffset < 0 {
		chinaOffset += 24
	}
	usaOffset := 12 + offset + 5
	log.Println("offset", offset, "chinaOffset:", chinaOffset, "usaOffset:", usaOffset)
	c := cron.New()
	c.AddFunc("0 0 "+strconv.Itoa(chinaOffset)+" * * *", func() {
		do(crawlerTargetChina)
	})
	c.AddFunc("0 0 "+strconv.Itoa(usaOffset)+" * * *", func() {
		do(crawlerTargetForgin)
	})
	c.Start()

	// Web服务
	if !rwn.C.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"tf": func(t time.Time) string {
			return t.Format("2006-01-02 15:04")
		},
	})
	r.Static("/static", "resource/static")
	r.LoadHTMLGlob("resource/template/*")
	r.GET("/", func(ctx *gin.Context) {
		var jobs []struct {
			Day  string
			Jobs []rwn.News
		}
		var news []rwn.News
		now := time.Now().AddDate(0, 0, -2)
		rwn.DB.Where("created_at > ?",
			fmt.Sprintf("%d-%d-%d 23:59:59", now.Year(), now.Month(), now.Day())).
			Order("created_at DESC").Find(&news)
		var currKey string
		var job struct {
			Day  string
			Jobs []rwn.News
		}
		for i := 0; i < len(news); i++ {
			key := news[i].CreatedAt.Format("2006-01-02")
			if key != currKey {
				currKey = key
				if len(job.Jobs) > 0 {
					jobs = append(jobs, job)
				}
				job = struct {
					Day  string
					Jobs []rwn.News
				}{
					Day:  currKey,
					Jobs: make([]rwn.News, 0),
				}
			}
			job.Jobs = append(job.Jobs, news[i])
		}
		jobs = append(jobs, job)
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"media":    rwn.Medias,
			"job":      jobs,
			"crawling": crawling,
			"version":  rwn.C.BuildVersion,
		})
	})
	r.Run()
}

func do(c []crawlers.Crawler) {
	if crawling {
		serverChan("「远程工作」抓取冲突了", "")
		return
	}
	crawling = true
	var errorMsg []byte
	var allNews []rwn.News
	var l sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(c))
	for i := 0; i < len(c); i++ {
		go func(i int) {
			news, err := c[i].FetchNews()
			if err != nil {
				errorMsg = append(errorMsg, ("- " + reflect.TypeOf(c[i]).String() + ":" + err.Error() + "\n")...)
			}
			err = c[i].FillContent(news)
			if err != nil {
				errorMsg = append(errorMsg, ("- " + reflect.TypeOf(c[i]).String() + ":" + err.Error() + "\n")...)
			}
			l.Lock()
			allNews = append(allNews, news...)
			l.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	now := time.Now()
	for i := 0; i < len(allNews); i++ {
		allNews[i].Title = strings.TrimSpace(allNews[i].Title)
		allNews[i].Hash = com.MD5(allNews[i].URL)
		allNews[i].CreatedAt = now
		rwn.DB.Save(allNews[i])
	}
	if len(errorMsg) > 0 {
		serverChan("「远程工作」抓取错误", string(errorMsg))
	}
	crawling = false
}

func serverChan(title, content string) {
	params := url.Values{
		"text": {title},
		"desp": {content},
	}
	http.PostForm("https://sc.ftqq.com/"+rwn.C.ServerChan+".send", params)
}
