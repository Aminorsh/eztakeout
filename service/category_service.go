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

func (s *CategoryService) Delete(id uint64) error {
	return s.DB.Delete(&model.Category{}, id).Error
}

func (s *CategoryService) Update(category *model.Category) error {
	return s.DB.Model(&model.Category{}).
		Where("id = ?", category.ID).
		Updates(map[string]any{
			"name": category.Name,
			"type": category.Type,
			"sort": category.Sort,
		}).Error
}

func (s *CategoryService) Page(page, pageSize int) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	err := s.DB.Model(&model.Category{}).
		Count(&total).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("sort asc").
		Find(&categories).Error

	return categories, total, err
}
