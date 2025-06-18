package service

import (
	"eztakeout/dto"
	"eztakeout/model"

	"gorm.io/gorm"
)

type SetmealService struct {
	DB *gorm.DB
}

type SetmealV0 struct {
	model.Setmeal
	CategoryName string `json:"category_name"`
}

type SetmealPageResult struct {
	Total    int64       `json:"total"`
	Records  []SetmealV0 `json:"records"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

func (s *SetmealService) Add(setmeal *model.Setmeal, dishes []model.SetmealDish) error {
	return s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(setmeal).Error; err != nil {
			return err
		}

		for i := range dishes {
			dishes[i].SetmealID = setmeal.ID
		}

		return tx.Create(&dishes).Error
	})
}

func (s *SetmealService) Page(query dto.SetmealPageQuery) (*SetmealPageResult, error) {
	var total int64
	var list []model.Setmeal

	tx := s.DB.Model(&model.Setmeal{}).Where("is_deleted = 0")
	if query.Name != "" {
		tx = tx.Where("name LIKE ?", "%"+query.Name+"%")
	}

	if err := tx.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (query.Page - 1) * query.PageSize
	if err := tx.Order("update_time desc").
		Offset(offset).
		Limit(query.PageSize).
		Find(&list).Error; err != nil {
		return nil, err
	}

	var result []SetmealV0
	for _, item := range list {
		var category model.Category
		s.DB.First(&category, item.CategoryID)

		result = append(result, SetmealV0{
			Setmeal:      item,
			CategoryName: category.Name,
		})
	}

	return &SetmealPageResult{
		Total:    total,
		Records:  result,
		Page:     query.Page,
		PageSize: query.PageSize,
	}, nil
}

func (s *SetmealService) UpdateStatus(id uint64, status int) error {
	return s.DB.Model(&model.Setmeal{}).
		Where("id = ?", id).
		Update("status", status).Error
}
