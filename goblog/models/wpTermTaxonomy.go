package models

type WpTermTaxonomy struct {
	TermTaxonomyId int64  `xorm:"not null pk autoincr BIGINT(20)"`
	TermId         int64  `xorm:"not null default 0 unique(term_id_taxonomy) BIGINT(20)"`
	Taxonomy       string `xorm:"not null default '' unique(term_id_taxonomy) index VARCHAR(32)"`
	Description    string `xorm:"not null LONGTEXT"`
	Parent         int64  `xorm:"not null default 0 BIGINT(20)"`
	Count          int64  `xorm:"not null default 0 BIGINT(20)"`
}
