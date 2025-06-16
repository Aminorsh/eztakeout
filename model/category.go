package model

import (
	"time"
)

type Category struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(32);not null" json:"name"`
	Type        int       `gorm:"type:int;not null" json:"type"` // 1: single dish, 2: set meal
	Sort        int       `gorm:"type:int;default:0" json:"sort"`
	Status      int       `gorm:"type:tinyint;default:1" json:"status"` // 1: active, 0: inactive
	CreatedTime time.Time `gorm:"autoCreateTime" json:"created_time"`
	UpdatedTime time.Time `gorm:"autoUpdateTime" json:"updated_time"`
}

// TableName overrides the default table name for Category model
func (Category) TableName() string {
	return "category"
}
