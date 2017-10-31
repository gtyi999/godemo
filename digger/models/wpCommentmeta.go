package models

type WpCommentmeta struct {
	MetaId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	CommentId int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}
