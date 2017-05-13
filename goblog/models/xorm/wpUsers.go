package xorm

import (
	"time"
)

type WpUsers struct {
	Id                int64     `xorm:"pk autoincr BIGINT(20)"`
	UserLogin         string    `xorm:"not null default '' index VARCHAR(60)"`
	UserPass          string    `xorm:"not null default '' VARCHAR(255)"`
	UserNicename      string    `xorm:"not null default '' index VARCHAR(50)"`
	UserEmail         string    `xorm:"not null default '' VARCHAR(100)"`
	UserUrl           string    `xorm:"not null default '' VARCHAR(100)"`
	UserRegistered    time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UserActivationKey string    `xorm:"not null default '' VARCHAR(255)"`
	UserStatus        int       `xorm:"not null default 0 INT(11)"`
	DisplayName       string    `xorm:"not null default '' VARCHAR(250)"`
}
