package seed

import (
	"log"
	"time"

	"github.com/omnlgy/RESTHARISGO/internal/models"
	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	departments := seedDepartments(db)
	positions := seedPositions(db)
	employees := seedEmployees(db, departments, positions)
	seedAttendances(db, employees)
	seedLeaves(db, employees)
	seedSalaries(db, employees)
	log.Println("Seed data inserted successfully")
}

func seedDepartments(db *gorm.DB) []models.Department {
	items := []models.Department{
		{Name: "IT", Code: "DEPT-IT"},
		{Name: "HR", Code: "DEPT-HR"},
		{Name: "Finance", Code: "DEPT-FIN"},
		{Name: "Marketing", Code: "DEPT-MKT"},
		{Name: "Operations", Code: "DEPT-OPS"},
	}
	for i := range items {
		db.FirstOrCreate(&items[i], models.Department{Code: items[i].Code})
	}
	return items
}

func seedPositions(db *gorm.DB) []models.Position {
	items := []models.Position{
		{Title: "Software Engineer", BaseSalary: 8_000_000},
		{Title: "Senior Software Engineer", BaseSalary: 15_000_000},
		{Title: "HR Specialist", BaseSalary: 5_000_000},
		{Title: "Finance Analyst", BaseSalary: 6_000_000},
		{Title: "Marketing Lead", BaseSalary: 7_000_000},
		{Title: "Operations Manager", BaseSalary: 10_000_000},
	}
	for i := range items {
		db.FirstOrCreate(&items[i], models.Position{Title: items[i].Title})
	}
	return items
}

func seedEmployees(db *gorm.DB, departments []models.Department, positions []models.Position) []models.Employee {
	items := []models.Employee{
		{NIK: "EMP-001", FullName: "Andi Pratama", Email: "andi@company.com", DepartmentID: departments[0].ID, PositionID: positions[0].ID, Status: "ACTIVE"},
		{NIK: "EMP-002", FullName: "Siti Rahayu", Email: "siti@company.com", DepartmentID: departments[0].ID, PositionID: positions[1].ID, Status: "ACTIVE"},
		{NIK: "EMP-003", FullName: "Budi Santoso", Email: "budi@company.com", DepartmentID: departments[1].ID, PositionID: positions[2].ID, Status: "ACTIVE"},
		{NIK: "EMP-004", FullName: "Dewi Lestari", Email: "dewi@company.com", DepartmentID: departments[2].ID, PositionID: positions[3].ID, Status: "ACTIVE"},
		{NIK: "EMP-005", FullName: "Rudi Hartono", Email: "rudi@company.com", DepartmentID: departments[3].ID, PositionID: positions[4].ID, Status: "SUSPENDED"},
		{NIK: "EMP-006", FullName: "Maya Indah", Email: "maya@company.com", DepartmentID: departments[4].ID, PositionID: positions[5].ID, Status: "ACTIVE"},
	}
	for i := range items {
		db.FirstOrCreate(&items[i], models.Employee{NIK: items[i].NIK})
	}
	return items
}

func seedAttendances(db *gorm.DB, employees []models.Employee) {
	// Clear existing attendances for clean seed
	db.Where("1 = 1").Delete(&models.Attendance{})

	today := time.Now().Truncate(24 * time.Hour)
	for _, emp := range employees {
		if emp.Status != "ACTIVE" {
			continue
		}
		for dayOffset := range 5 {
			date := today.AddDate(0, 0, -dayOffset)
			dateStr := date.Format("2006-01-02")
			if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
				continue
			}

			status := "PRESENT"
			checkIn := "08:00"
			checkOut := "17:00"

			switch emp.NIK {
			case "EMP-002":
				checkIn = "08:45"
				status = "LATE"
			case "EMP-006":
				checkIn = "09:10"
				status = "LATE"
				checkOut = "17:30"
			}

			db.FirstOrCreate(&models.Attendance{}, map[string]interface{}{
				"employee_id": emp.ID,
				"date":        dateStr,
			}, &models.Attendance{
				EmployeeID: emp.ID,
				Date:       dateStr,
				CheckIn:    checkIn,
				CheckOut:   checkOut,
				Status:     status,
			})
		}
	}
}

func seedLeaves(db *gorm.DB, employees []models.Employee) {
	// Clear existing leaves for clean seed
	db.Where("1 = 1").Delete(&models.Leave{})

	// Approved leave for employee 1
	db.FirstOrCreate(&models.Leave{}, models.Leave{
		EmployeeID: employees[0].ID,
		StartDate:  "2026-07-01",
		EndDate:    "2026-07-03",
		Reason:     "Cuti tahunan",
		Status:     "APPROVED",
	})

	// Pending leave for employee 3
	db.FirstOrCreate(&models.Leave{}, models.Leave{
		EmployeeID: employees[2].ID,
		StartDate:  "2026-08-10",
		EndDate:    "2026-08-12",
		Reason:     "Acara keluarga",
		Status:     "PENDING",
	})

	// Rejected leave for employee 5
	db.FirstOrCreate(&models.Leave{}, models.Leave{
		EmployeeID: employees[4].ID,
		StartDate:  "2026-06-15",
		EndDate:    "2026-06-16",
		Reason:     "Izin pribadi",
		Status:     "REJECTED",
	})
}

func seedSalaries(db *gorm.DB, employees []models.Employee) {
	// Clear existing salaries for clean seed
	db.Where("1 = 1").Delete(&models.Salary{})

	for _, emp := range employees {
		db.FirstOrCreate(&models.Salary{}, models.Salary{
			EmployeeID:  emp.ID,
			Period:      "2026-06",
			BasicSalary: 8_000_000,
			Allowance:   500_000,
			Deductions:  100_000,
		})
	}
}
