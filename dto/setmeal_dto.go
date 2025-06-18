package dto

type SetmealDishDTO struct {
	DishID uint64  `json:"dish_id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Copies int     `json:"copies"`
	Sort   int     `json:"sort"`
}

type SetmealAddDTO struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Image       string           `json:"image"`
	Price       float64          `json:"price"`
	CategoryID  uint64           `json:"category_id"`
	Dishes      []SetmealDishDTO `json:"dishes"`
}

type SetmealPageQuery struct {
	Name     string `json:"name"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type SetmealStatusDTO struct {
	ID     uint64 `json:"id" binding:"required"`
	Status *int   `json:"status" binding:"required,oneof=0 1"`
}
