package dto

type DishPageQueryDTO struct {
	Name     string `form:"name"`
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"pageSize" binding:"required"`
}
