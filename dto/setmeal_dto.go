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
