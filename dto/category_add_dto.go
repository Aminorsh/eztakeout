package dto

type CategoryAddDTO struct {
	Name string `json:"name" binding:"required"`
	Type int    `json:"type" binding:"required"`
	Sort int    `json:"sort"`
}
