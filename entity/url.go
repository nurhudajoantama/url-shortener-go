package entity

import "time"

type Url struct {
	ID        uint64 `gorm:"primary_key:auto_increment"`
	ShortUrl  string `gorm:"type:varchar(255),unique"`
	LongUrl   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
