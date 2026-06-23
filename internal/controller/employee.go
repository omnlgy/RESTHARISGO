package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/RESTHARISGO/internal/models"
	"github.com/omnlgy/RESTHARISGO/internal/repository"
	"github.com/omnlgy/RESTHARISGO/internal/service"
)

type CreateEmployeeRequest struct {
	NIK          string `json:"nik" binding:"required"`
	FullName     string `json:"fullName" binding:"required"`
	Email        string `json:"email" binding:"required"`
	DepartmentID uint   `json:"departmentId" binding:"required"`
	PositionID   uint   `json:"positionId" binding:"required"`
	Status       string `json:"status" binding:"required,oneof=ACTIVE SUSPENDED TERMINATED"`
}

type EmployeeController struct {
	service *service.EmployeeService
}

func NewEmployeeController(service *service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
	}
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	var body CreateEmployeeRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	employee := &models.Employee{
		NIK:          body.NIK,
		FullName:     body.FullName,
		Email:        body.Email,
		DepartmentID: body.DepartmentID,
		PositionID:   body.PositionID,
		Status:       body.Status,
	}

	if _, err := c.service.Add(employee); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ctx.JSON(400, gin.H{"error": "Department or Position not found"})
			return
		}
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Employee created successfully",
		"data":    employee,
	})
}

func (c *EmployeeController) GetEmployees(ctx *gin.Context) {
	departmentID, _ := strconv.ParseUint(ctx.Query("departmentId"), 10, 64)
	positionID, _ := strconv.ParseUint(ctx.Query("positionId"), 10, 64)

	filter := models.FilterEmployee{
		Name:         ctx.Query("search"),
		Status:       ctx.Query("status"),
		DepartmentID: uint(departmentID),
		PositionID:   uint(positionID),
	}
	employees, err := c.service.GetEmployees(filter)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Employees retrieved successfully",
		"data":    employees,
		"query":   filter,
	})
}
