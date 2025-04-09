package models

import "time"

type Posts struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Title        string    `gorm:"type:varchar(200);not null" json:"title"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	Category     string    `gorm:"type:varchar(100);not null" json:"category"`
	Created_date time.Time `gorm:"type:timestamp" json:"created_date"`
	Updated_date time.Time `gorm:"type:timestamp" json:"updated_date"`
	Status       string    `gorm:"type:varchar(100);not null;check:status IN ('publish', 'draft', 'thrash')" json:"status"`
}
