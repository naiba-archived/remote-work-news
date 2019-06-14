package main

import (
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
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

	// test code
	// x := &crawlers.StackOverFlowCrawler{}
	// log.Println(x.FetchNews())
	// os.Exit(0)

	var crawler = []crawlers.Crawler{
		&crawlers.StackOverFlowCrawler{},
		&crawlers.VueJobsCrawler{},
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
	offset = 0
	chinaOffset := 12 + offset - 8
	if chinaOffset < 0 {
		chinaOffset += 24
	}
	c := cron.New()
	c.AddFunc("0 0 "+strconv.Itoa(chinaOffset)+" * * *", func() {
		do(crawler)
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
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()
		},
	})
	r.Static("/static", "resource/static")
	r.LoadHTMLGlob("resource/template/*")
	hourDiff := chinaOffset - 13
	dayDiff := -2
	if hourDiff < 0 {
		hourDiff += 24
	}
	println("dayDiff", dayDiff)
	println("hourDiff", hourDiff)
	println("chinaOffset", chinaOffset)
	r.GET("/", func(ctx *gin.Context) {
		var jobs []struct {
			Day  string
			Jobs []rwn.News
		}
		var news []rwn.News
		rwn.DB.Order("created_at DESC").Limit(50).Find(&news)
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
		crawlers.ClearSpace(allNews)
		allNews[i].Hash = com.MD5(allNews[i].URL)
		allNews[i].CreatedAt = now
		rwn.DB.Save(allNews[i])
	}
	if len(errorMsg) > 0 {
		serverChan("「远程工作」抓取错误", string(errorMsg))
	} else {
		serverChan("「远程工作」抓取完成", time.Now().String())
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
