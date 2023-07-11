package routes

import (
	"github.com/gin-gonic/gin"
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"net/http"
	"strconv"
)

func GetEmplpoyee(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}
	employees := []models.Employee{}
	db.Find(&employees)

	employeeResponse := []models.EmployeeResponse{}

	for _, p := range employees {
		position := models.Position{}
		db.First(&position, "id = ?", p.PositionId)
		responsePosition := models.PositionResponseInnerJoin{
			ID:   position.ID,
			Name: position.Name,
			Code: position.Code,
		}
		employee := models.EmployeeResponse{
			ID:         p.ID,
			Name:       p.Name,
			Address:    p.Address,
			Email:      p.Email,
			PositionId: p.PositionId,
			Position:   responsePosition,
		}

		employeeResponse = append(employeeResponse, employee)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get Position",
		"data":    employeeResponse,
	})
}

func PostEmployee(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	var newEmployee models.EmployeeRequest

	err = c.ShouldBindJSON(&newEmployee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}

	employee := models.Employee{
		Name:       newEmployee.Name,
		Address:    newEmployee.Address,
		Email:      newEmployee.Email,
		PositionId: newEmployee.PositionId,
	}
	err = db.Create(&employee).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    newEmployee,
	})
}

func PutEmployee(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	stringId := c.Param("id")

	id, _ := strconv.Atoi(stringId)

	var employee models.Employee
	err = c.ShouldBindJSON(&employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Message": err.Error(),
		})
		return
	}

	//config.DB.First(&department, "id = ?", id)
	result := db.Model(models.Employee{}).Debug().Where("id = ?", id).Updates(&employee)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    employee,
	})
}

func DeleteEmployee(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	stringId := c.Param("id")

	id, _ := strconv.Atoi(stringId)

	var employee models.Employee
	resDB := db.Delete(&employee, "id = ?", id)

	if resDB.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"MSG": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Data berhasil terhapus",
	})
}
