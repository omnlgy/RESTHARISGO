package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/omnlgy/RESTHARISGO/internal/controller"
	"github.com/omnlgy/RESTHARISGO/internal/repository"
	"github.com/omnlgy/RESTHARISGO/internal/router"
	"github.com/omnlgy/RESTHARISGO/internal/service"
	"github.com/omnlgy/RESTHARISGO/internal/utils"
	"gorm.io/gorm"
)

func main() {
	// Load .env (optional; won't error if file is missing)
	_ = godotenv.Load()

	db, err := utils.InitDB()

	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
		return
	}

	// Gin server
	server := gin.Default()

	// Initialize routes
	initRouter(server, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := server.Run(":" + port); err != nil {
		fmt.Println("Failed to start server")
	}
}

func initRouter(server *gin.Engine, db *gorm.DB) {
	// Department
	repoDep := repository.NewDepartmentRepository(db)
	svcDep := service.NewDepartmentService(repoDep)
	ctrlDep := controller.NewDepartmentController(svcDep)
	router.DepartmentRoutes(server, ctrlDep)

	// Position
	repoPos := repository.NewPositionRepository(db)
	svcPos := service.NewPositionService(repoPos)
	ctrlPos := controller.NewPositionController(svcPos)
	router.PositionRoutes(server, ctrlPos)

	// Employee
	repoEmp := repository.NewEmployeeRepository(db)
	svcEmp := service.NewEmployeeService(repoEmp, svcDep, svcPos)
	ctrlEmp := controller.NewEmployeeController(svcEmp)
	router.EmployeeRoutes(server, ctrlEmp)
}
