package rwn

// Media 抓取平台
type Media struct {
	Name string
	Desc string
	Link string
}

// Medias 所有平台列表
var Medias map[int]Media

func init() {
	Medias = map[int]Media{
		1: Media{
			Name: "一早一晚",
			Desc: "一早一晚旨在帮助更多人走上「只工作，不上班」的自由工作之路，我们是一个「分布式组织」，通过分享及行动带来积极的影响，相信点滴的力量能改变潮水的方向。",
			Link: "https://yizaoyiwan.com",
		},
		2: Media{
			Name: "RubyChina",
			Desc: "由众多爱好者共同维护的 Ruby 中文社区，使用 Homeland 构建，并采用 Docker 部署。",
			Link: "https://ruby-china.org",
		},
	}
}
