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
		3: Media{
			Name: "LearnKu",
			Desc: "LearnKu 是终身学习者的编程知识社区：保持好奇，刻意练习，每日精进",
			Link: "https://learnku.com",
		},
		4: Media{
			Name: "SegmentFault",
			Desc: "SegmentFault 思否 为开发者提供问答、学习与交流编程知识的平台,创造属于开发者的时代!",
			Link: "https://segmentfault.com",
		},
		5: Media{
			Name: "StackOverFlow",
			Desc: "Stack Overflow是一个程序设计领域的问答网站，隶属Stack Exchange Network。网站允许注册用户提出或回答问题，还可对已有问题或答案加分、扣分或进行修改，条件是用户达到一定的“声望值”。",
			Link: "https://stackoverflow.com",
		},
		6: Media{
			Name: "远程.work",
			Desc: "远程.work是远程工作招聘网站,帮助企业和团队找到远程工作员工,为远程工作者寻找远程工作。",
			Link: "https://yuancheng.work",
		},
		7: Media{
			Name: "VueJobs",
			Desc: "VueJobs has been developed by Vue.js and Laravel enthusiasts whose aim is to contribute to the growing community by helping companies find Vue.js talent around the world. When you register with Vue.js jobs, companies will be able to list jobs and developers will have the chance to apply for those jobs. ",
			Link: "https://vuejobs.com",
		},
	}
}
