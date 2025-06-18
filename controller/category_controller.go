package controller

import (
	"eztakeout/dto"
	"eztakeout/model"
	"eztakeout/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	Service *service.CategoryService
}

func (c *CategoryController) Add(ctx *gin.Context) {
	var req dto.CategoryAddDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	category := model.Category{
		Name:   req.Name,
		Type:   req.Type,
		Sort:   req.Sort,
		Status: 1,
	}

	if err := c.Service.Add(&category); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to add category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Category added successfully"})
}

func (c *CategoryController) List(ctx *gin.Context) {
	typeStr := ctx.Query("type")
	if typeStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Type parameter is required"})
		return
	}

	var categoryType int
	if typeStr == "1" {
		categoryType = 1 // Single dish
	} else if typeStr == "2" {
		categoryType = 2 // Set meal
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Invalid type parameter"})
		return
	}

	categories, err := c.Service.List(categoryType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to retrieve categories"})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "ID parameter is required"})
		return
	}

	var id uint64
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Invalid ID parameter"})
		return
	}

	if err := c.Service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "Failed to delete category"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "message": "Category deleted successfully"})
}
