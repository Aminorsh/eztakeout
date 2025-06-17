package service

import (
	"eztakeout/model"

	"gorm.io/gorm"
)

type CategoryService struct {
	DB *gorm.DB
}

func (s *CategoryService) Add(category *model.Category) error {
	return s.DB.Create(category).Error
}

func (s *CategoryService) List(categoryType int) ([]model.Category, error) {
	var categories []model.Category
	err := s.DB.Where("type = ?", categoryType).Order("sort asc").Find(&categories).Error
	return categories, err
}
