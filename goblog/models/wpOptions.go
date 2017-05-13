package models

type WpOptions struct {
	OptionId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	OptionName  string `xorm:"not null default '' unique VARCHAR(191)"`
	OptionValue string `xorm:"not null LONGTEXT"`
	Autoload    string `xorm:"not null default 'yes' VARCHAR(20)"`
}
