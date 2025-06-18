package router

import (
	"eztakeout/controller"
	"eztakeout/middleware"
	"eztakeout/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Initialize services
	empService := &service.EmployeeService{DB: db}
	catService := &service.CategoryService{DB: db}
	dishService := &service.DishService{DB: db}
	setService := &service.SetmealService{DB: db}

	// Initialize controllers
	empController := &controller.EmployeeController{Service: empService}
	catController := &controller.CategoryController{Service: catService}
	dishController := &controller.DishController{Service: dishService}
	setController := &controller.SetmealController{Service: setService}

	// Define routes
	r.POST("/login", empController.Login)
	// r.POST("/categories", catController.Add)
	// r.GET("/categories", catController.List)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/categories", catController.Add)
		authorized.GET("/categories", catController.List)
		authorized.DELETE("/categories/:id", catController.Delete)
		authorized.PUT("/categories", catController.Update)

		authorized.GET("/categories/page", catController.Page)

		authorized.POST("/dishes", dishController.Add)
		authorized.GET("/dishes/page", dishController.Page)

		authorized.PUT("/dishes", dishController.Update)
		authorized.PUT("/dishes/:id/status", dishController.UpdateStatus)
		authorized.DELETE("/dishes", dishController.Delete)
		authorized.GET("/dishes/list", dishController.ListByCategory)
		authorized.DELETE("/dishes/:id", dishController.DeleteByID)

		authorized.POST("/setmeals", setController.Add)
		authorized.GET("/setmeals/page", setController.Page)
		authorized.PUT("/setmeals/status", setController.UpdateStatus)
		authorized.DELETE("/setmeals", setController.Delete)
	}

	return r
}
