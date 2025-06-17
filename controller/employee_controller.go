package controller

import (
	"eztakeout/dto"
	"eztakeout/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Service *service.EmployeeService
}

func (c *EmployeeController) Login(ctx *gin.Context) {
	var loginDto dto.EmployeeLoginDTO
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "Parameter error"})
		return
	}

	emp, err := c.Service.Login(loginDto.Username, loginDto.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": err.Error()})
		return
	}

	ctx.SetCookie("session_id", "logged_in", 60, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data": gin.H{
			"id":   emp.ID,
			"name": emp.Name,
		},
	})
}
