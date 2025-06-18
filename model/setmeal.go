package model

import "time"

type Setmeal struct {
	ID          uint64    `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(64);not null" json:"name"`
	Description string    `gorm:"type:varchar(255)" json:"description"`
	Image       string    `gorm:"type:varchar(255)" json:"image"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	CategoryID  uint64    `gorm:"not null" json:"category_id"`
	Status      int       `gorm:"type:tinyint;default:1" json:"status"`
	IsDeleted   int       `gorm:"type:tinyint;default:0" json:"is_deleted"`
	CreateTime  time.Time `gorm:"autoCreateTime" json:"create_time"`
	UpdateTime  time.Time `gorm:"autoUpdateTime" json:"update_time"`
}

type SetmealDish struct {
	ID        uint64  `gorm:"primaryKey" json:"id"`
	SetmealID uint64  `gorm:"not null" json:"setmeal_id"`
	DishID    uint64  `gorm:"not null" json:"dish_id"`
	Name      string  `gorm:"type:varchar(64);not null" json:"name"`
	Price     float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Copies    int     `gorm:"type:int;default:1" json:"copies"`
	Sort      int     `gorm:"type:int;default:0" json:"sort"`
}

func (Setmeal) TableName() string {
	return "setmeal"
}
func (SetmealDish) TableName() string {
	return "setmeal_dish"
}
