package dto

type DishUpdateDTO struct {
	ID          uint64  `json:"id" binding:"required"`
	Name        string  `json:"name"`
	CategoryID  uint64  `json:"categoryId"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
}
