package dto

type DishAddDTO struct {
	Name        string  `json:"name" binding:"required"`
	CategoryID  uint64  `json:"category_id" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}
