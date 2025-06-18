package controller

import (
	"eztakeout/dto"
	"eztakeout/model"
	"eztakeout/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DishController struct {
	Service *service.DishService
}

func (c *DishController) Add(ctx *gin.Context) {
	var req dto.DishAddDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	dish := model.Dish{
		Name:        req.Name,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		Image:       req.Image,
		Description: req.Description,
		Status:      1,
	}

	if err := c.Service.Add(&dish); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to add dish"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Dish added successfully", "data": dish})
}

func (c *DishController) Page(ctx *gin.Context) {
	var req dto.DishPageQueryDTO
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	dishes, total, err := c.Service.PageList(req.Name, req.Page, req.PageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to retrieve dishes"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"total": total,
			"list":  dishes,
		},
	})
}
