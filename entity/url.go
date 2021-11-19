package entity

import "time"

type Url struct {
	ID        uint64 `gorm:"primary_key:auto_increment"`
	ShortUrl  string `gorm:"type:varchar(255); collate:utf8mb4_bin; index:short_url_index,unique,type:btree; not null"`
	LongUrl   string `gorm:"type:text;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
