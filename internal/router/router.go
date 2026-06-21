package router

import (
	"github.com/gin-gonic/gin"
	"github.com/omnlgy/RESTHARISGO/internal/controller"
)

func DepartmentRoutes(router *gin.Engine, controller *controller.DepartmentController) {
	apiDepartments := router.Group("/api/departments")

	apiDepartments.GET("/", controller.GetDepartments)
	apiDepartments.POST("/", controller.CreateDepartment)
	apiDepartments.PUT("/:id", controller.UpdateDepartment)
	apiDepartments.DELETE("/:id", controller.DeleteDepartment)
}
