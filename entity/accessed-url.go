package entity

import "time"

type AccessedUrl struct {
	ID        uint64
	Url       Url
	AccesedAt time.Time
}
