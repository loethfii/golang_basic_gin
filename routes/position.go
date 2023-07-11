package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"net/http"
	"strconv"
)

func GetPosition(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}
	positions := []models.Position{}
	db.Find(&positions)

	positionsResponse := []models.PositionResponse{}

	for _, p := range positions {
		department := models.Department{}
		db.First(&department, "id = ?", p.DepartmentId)
		responseDepartment := models.DepartmentResponse{
			ID:   department.ID,
			Name: department.Name,
			Code: department.Code,
		}
		pos := models.PositionResponse{
			ID:           p.ID,
			Name:         p.Name,
			Code:         p.Code,
			DepartmentId: p.DepartmentId,
			Departments:  responseDepartment,
		}

		positionsResponse = append(positionsResponse, pos)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get Position",
		"data":    positionsResponse,
	})
}

func PostPosition(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	var newPosition models.PositionRequest

	err = c.ShouldBindJSON(&newPosition)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	position := models.Position{
		Name:         newPosition.Name,
		Code:         newPosition.Code,
		DepartmentId: newPosition.DepartmentId,
	}
	err = db.Create(&position).Error
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    newPosition,
	})
}

func PutPosition(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	stringId := c.Param("id")

	id, _ := strconv.Atoi(stringId)

	var position models.Position
	err = c.BindJSON(&position)
	if err != nil {
		panic(err)
	}

	//config.DB.First(&department, "id = ?", id)
	result := db.Model(models.Position{}).Debug().Where("id = ?", id).Updates(&position)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Position not found"})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    position,
	})
}

func DeletePosition(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	stringId := c.Param("id")

	id, _ := strconv.Atoi(stringId)

	var position models.Position
	resDB := db.Delete(&position, "id = ?", id)

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
