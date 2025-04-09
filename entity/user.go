package entity

type Posts struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Title        string `gorm:"type:varchar(200);not null" json:"title"`
	Content      string `gorm:"type:text;not null" json:"content"`
	Category     string `gorm:"type:varchar(100);not null" json:"category"`
	Created_date string `gorm:"type:datetime;default:null" json:"created_date"`
	Updated_date string `gorm:"type:datetime;default:null" json:"updated_date"`
	Status       string `gorm:"type:varchar(100);not null;check:status IN ('Publish', 'Draft', 'Thrash')" json:"status"`
}
