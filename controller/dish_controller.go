package controller

import (
	"eztakeout/dto"
	"eztakeout/model"
	"eztakeout/service"
	"fmt"
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

func (c *DishController) Update(ctx *gin.Context) {
	var req dto.DishUpdateDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	dish := model.Dish{
		ID:          req.ID,
		Name:        req.Name,
		CategoryID:  req.CategoryID,
		Price:       req.Price,
		Image:       req.Image,
		Description: req.Description,
	}

	if err := c.Service.Update(&dish); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to update dish"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Dish updated successfully", "data": dish})
}

func (c *DishController) UpdateStatus(ctx *gin.Context) {
	idStr := ctx.Param("id")
	statusStr := ctx.Query("status")

	var id uint64
	var status int
	_, err1 := fmt.Sscan(idStr, &id)
	_, err2 := fmt.Sscan(statusStr, &status)
	if err1 != nil || err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Invalid parameters"})
		return
	}

	if err := c.Service.UpdateStatus(id, status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to update dish status"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Dish status updated successfully"})
}

func (c *DishController) Delete(ctx *gin.Context) {
	var req struct {
		IDs []uint64 `json:"ids"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil || len(req.IDs) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	if err := c.Service.Delete(req.IDs); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to delete dishes"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Dishes deleted successfully"})
}

func (c *DishController) ListByCategory(ctx *gin.Context) {
	categoryIDStr := ctx.Query("category_id")
	var categoryID uint64
	if _, err := fmt.Sscan(categoryIDStr, &categoryID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Invalid category ID"})
		return
	}

	dishes, err := c.Service.ListByCategory(categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to retrieve"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "success", "data": dishes})
}
