package dto

type CategoryUpdateDTO struct {
	ID   uint64 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Type int    `json:"type" binding:"required"`
	Sort int    `json:"sort"`
}
