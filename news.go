package rwn

import "time"

// News 招聘信息
type News struct {
	Hash       string
	MediaID    uint
	Title      string
	URL        string
	Content    string
	Pusher     string
	PusherLink string
	CreatedAt  time.Time
}
