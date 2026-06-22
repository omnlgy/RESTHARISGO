package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omnlgy/RESTHARISGO/internal/models"
	"github.com/omnlgy/RESTHARISGO/internal/service"
)

type CreateDepartmentRequest struct {
	Title string `json:"title" binding:"required"`
	BaseSalary float64 `json:"baseSalary" binding:"required"`
}

type PositionController struct {
	service *service.PositionService
}

func NewPositionController(service *service.PositionService) *PositionController {
	return &PositionController{
		service: service,
	}
}

func (c *PositionController) GetPositions(ctx *gin.Context) {
	positions, err := c.service.GetPositions()

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Positions retrieved successfully",
		"data":    positions,
	})
}

func (c *PositionController) CreatePosition(ctx *gin.Context) {
	var body CreateDepartmentRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	position := &models.Position{
		Title: body.title,
		BaseSalary: body.baseSalary,
	}

	createdDepartment, err := c.service.CreatePosition(position)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Department created successfully",
		"data":    createdDepartment,
	})
}

func (c *PositionController) UpdatePosition(ctx *gin.Context) {
	var body CreateDepartmentRequest
	var PositionID uint64

	PositionID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid position ID"})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	position := &models.Position{
		ID:   uint(PositionID),
		Title: body.title,
		BaseSalary: body.baseSalary,
	}

	updatedPosition, err := c.service.UpdatePosition(position)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Position updated successfully",
		"data":    updatedPosition,
	})
}

func (c *PositionController) DeletePosition(ctx *gin.Context) {
	positionID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid department ID"})
		return
	}

	err = c.service.DeletePosition(uint(positionID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Posetion deleted successfully",
	})
}
