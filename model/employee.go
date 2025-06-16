package model

import (
	"time"
)

type Employee struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Username    string    `gorm:"type:varchar(32);unique;not null" json:"username"`
	Password    string    `gorm:"type:varchar(64);not null" json:"password"`
	Name        string    `gorm:"type:varchar(32);not null" json:"name"`
	Phone       string    `gorm:"type:varchar(20);not null" json:"phone"`
	Status      int       `gorm:"type:tinyint;default:1" json:"status"` // 1: active, 0: inactive
	CreatedTime time.Time `gorm:"autoCreateTime" json:"created_time"`
	UpdatedTime time.Time `gorm:"autoUpdateTime" json:"updated_time"`
}

// TableName overrides the default table name for Employee model
func (Employee) TableName() string {
	return "employee"
}
