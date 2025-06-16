package service

import (
	"errors"
	"eztakeout/model"

	"gorm.io/gorm"
)

type EmployeeService struct {
	DB *gorm.DB
}

func (s *EmployeeService) Login(username, password string) (*model.Employee, error) {
	var emp model.Employee
	if err := s.DB.Where("username = ?", username).First(&emp).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if emp.Password != password {
		return nil, errors.New("incorrect password")
	}

	if emp.Status == 0 {
		return nil, errors.New("user is inactive")
	}

	return &emp, nil
}
