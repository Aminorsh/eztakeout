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

func (c *SetmealController) Page(ctx *gin.Context) {
	var query dto.SetmealPageQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	result, err := c.Service.Page(query)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to fetch setmeals"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    result,
	})
}

func (c *SetmealController) UpdateStatus(ctx *gin.Context) {
	var req dto.SetmealStatusDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error", "detail": err.Error()})
		return
	}

	if err := c.Service.UpdateStatus(req.ID, *req.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to update status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Status updated successfully"})
}

func (c *SetmealController) Delete(ctx *gin.Context) {
	var ids []uint64
	if err := ctx.ShouldBindJSON(&ids); err != nil || len(ids) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	if err := c.Service.DeleteByIDs(ids); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to delete setmeals"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Setmeals deleted successfully"})
}
