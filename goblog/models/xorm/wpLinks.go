package xorm

import (
	"time"
)

type WpLinks struct {
	LinkId          int64     `xorm:"not null pk autoincr BIGINT(20)"`
	LinkUrl         string    `xorm:"not null default '' VARCHAR(255)"`
	LinkName        string    `xorm:"not null default '' VARCHAR(255)"`
	LinkImage       string    `xorm:"not null default '' VARCHAR(255)"`
	LinkTarget      string    `xorm:"not null default '' VARCHAR(25)"`
	LinkDescription string    `xorm:"not null default '' VARCHAR(255)"`
	LinkVisible     string    `xorm:"not null default 'Y' index VARCHAR(20)"`
	LinkOwner       int64     `xorm:"not null default 1 BIGINT(20)"`
	LinkRating      int       `xorm:"not null default 0 INT(11)"`
	LinkUpdated     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	LinkRel         string    `xorm:"not null default '' VARCHAR(255)"`
	LinkNotes       string    `xorm:"not null MEDIUMTEXT"`
	LinkRss         string    `xorm:"not null default '' VARCHAR(255)"`
}
