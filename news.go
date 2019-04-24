package rwn

import "time"

// News 招聘信息
type News struct {
	Hash       string `gorm:"unique_index"`
	MediaID    uint
	Title      string
	URL        string
	Content    string
	Pusher     string
	PusherLink string
	PublishedAt  time.Time
	CreatedAt time.Time
}
