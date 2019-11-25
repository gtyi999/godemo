package models

type WpUsermeta struct {
	UmetaId   int64  `xorm:"not null pk autoincr BIGINT(20)"`
	UserId    int64  `xorm:"not null default 0 index BIGINT(20)"`
	MetaKey   string `xorm:"index VARCHAR(255)"`
	MetaValue string `xorm:"LONGTEXT"`
}
