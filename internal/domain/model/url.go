package model

import "time"

type URL struct {
	ID 		  	int64      `gorm:"id"`
	Code 	  	string     `gorm:"code"`
	OriginalURL string     `gorm:"original_url"`
	Visits   	int64      `gorm:"visits"`
	LastVisit   time.Time  `gorm:"last_visit"`
	CreatedAt   time.Time  `gorm:"created_at"`
}