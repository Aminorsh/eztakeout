package model

import "time"

type Dish struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(64);not null" json:"name"`
	CategoryID  uint64    `gorm:"not null" json:"category_id"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Status      int       `gorm:"type:tinyint;default:1" json:"status"`
	Create_time time.Time `gorm:"autoCreateTime" json:"create_time"`
	Update_time time.Time `gorm:"autoUpdateTime" json:"update_time"`
	IsDeleted   int       `gorm:"type:tinyint;default:0" json:"is_deleted"`
}

func (Dish) TableName() string {
	return "dish"
}
