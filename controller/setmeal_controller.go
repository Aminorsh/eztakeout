package controller

import (
	"eztakeout/dto"
	"eztakeout/model"
	"eztakeout/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SetmealController struct {
	Service *service.SetmealService
}

func (c *SetmealController) Add(ctx *gin.Context) {
	var req dto.SetmealAddDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	setmeal := model.Setmeal{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.Image,
		Price:       req.Price,
		CategoryID:  req.CategoryID,
		Status:      1,
	}

	var dishes []model.SetmealDish
	for _, d := range req.Dishes {
		dishes = append(dishes, model.SetmealDish{
			Name:   d.Name,
			Price:  d.Price,
			Copies: d.Copies,
			Sort:   d.Sort,
			DishID: d.DishID,
		})
	}

	if err := c.Service.Add(&setmeal, dishes); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to add setmeal"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Setmeal added successfully"})
}
