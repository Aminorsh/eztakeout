package service

import (
	"eztakeout/model"

	"gorm.io/gorm"
)

type DishService struct {
	DB *gorm.DB
}

func (s *DishService) Add(dish *model.Dish) error {
	return s.DB.Create(dish).Error
}

func (s *DishService) PageList(name string, page int, pageSize int) ([]model.Dish, int64, error) {
	var dishes []model.Dish
	var total int64

	query := s.DB.Model(&model.Dish{})

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("update_time desc").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&dishes).Error

	return dishes, total, err
}
