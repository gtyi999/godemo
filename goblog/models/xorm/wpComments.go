package xorm

import (
	"time"
)

type WpComments struct {
	CommentId          int64     `xorm:"not null pk autoincr BIGINT(20)"`
	CommentPostId      int64     `xorm:"not null default 0 index BIGINT(20)"`
	CommentAuthor      string    `xorm:"not null TINYTEXT"`
	CommentAuthorEmail string    `xorm:"not null default '' index VARCHAR(100)"`
	CommentAuthorUrl   string    `xorm:"not null default '' VARCHAR(200)"`
	CommentAuthorIp    string    `xorm:"not null default '' VARCHAR(100)"`
	CommentDate        time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	CommentDateGmt     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' index(comment_approved_date_gmt) index TIMESTAMP"`
	CommentContent     string    `xorm:"not null TEXT"`
	CommentKarma       int       `xorm:"not null default 0 INT(11)"`
	CommentApproved    string    `xorm:"not null default '1' index(comment_approved_date_gmt) VARCHAR(20)"`
	CommentAgent       string    `xorm:"not null default '' VARCHAR(255)"`
	CommentType        string    `xorm:"not null default '' VARCHAR(20)"`
	CommentParent      int64     `xorm:"not null default 0 index BIGINT(20)"`
	UserId             int64     `xorm:"not null default 0 BIGINT(20)"`
}
