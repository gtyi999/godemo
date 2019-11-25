package models

import (
	"time"
)

type WpPosts struct {
	Id            int64     `xorm:"pk autoincr index(type_status_date) BIGINT(20)"`
	PostAuthor    int64     `xorm:"not null default 0 index BIGINT(20)"`
	PostDate      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' index(type_status_date) TIMESTAMP"`
	PostContent   string    `xorm:"not null LONGTEXT"`
	PostTitle     string    `xorm:"not null TEXT"`
	PostExcerpt   string    `xorm:"not null TEXT"`
	PostStatus    string    `xorm:"not null default 'publish' index(type_status_date) VARCHAR(20)"`
	CommentStatus string    `xorm:"not null default 'open' VARCHAR(20)"`
	PostName      string    `xorm:"not null default '' index VARCHAR(200)"`
	PostHeadimg   string    `xorm:"not null VARCHAR(45)"`
	PostModified  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	Guid          string    `xorm:"not null default '' VARCHAR(255)"`
	PostType      string    `xorm:"not null default 'post' index(type_status_date) VARCHAR(20)"`
	CommentCount  int64     `xorm:"not null default 0 BIGINT(20)"`
	PostParent    int64     `xorm:"not null default 0 BIGINT(11)"`
}
