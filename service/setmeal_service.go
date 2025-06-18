package service

import (
	"eztakeout/model"

	"gorm.io/gorm"
)

type SetmealService struct {
	DB *gorm.DB
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
