package routes

import (
	"github.com/gin-gonic/gin"
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func GetDepartment(c *gin.Context) {
	var db, err = config.InitDB()
	if err != nil {
		panic(err.Error())
	}
	departments := []models.Department{}
	//config.DB.Find(&departments)
	//db.Find(&departments)
	db.Preload(clause.Associations).Find(&departments)

	GetDepatementResponses := []models.DepartmentResponse{}

	for _, d := range departments {
		resPositions := []models.PositionResponseInnerJoin{}
		for _, p := range d.Posistions {
			pos := models.PositionResponseInnerJoin{
				ID:   p.ID,
				Name: p.Name,
				Code: p.Code,
			}
			resPositions = append(resPositions, pos)
		}

		dept := models.DepartmentResponse{
			ID:         d.ID,
			Name:       d.Name,
			Code:       d.Code,
			Posistions: resPositions,
		}

		GetDepatementResponses = append(GetDepatementResponses, dept)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get Departments",
		"data":    GetDepatementResponses,
	})
}

func PostDepartment(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	var newDepartment models.DepartmentRequest

	err = c.BindJSON(&newDepartment)
	if err != nil {
		panic(err)
	}
	department := models.Department{
		Name: newDepartment.Name,
		Code: newDepartment.Code,
	}
	err = db.Create(&department).Error
	if err != nil {
		panic(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    newDepartment,
	})
}

func PutDepartment(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	stringId := c.Param("id")

	id, _ := strconv.Atoi(stringId)

	var department models.Department
	err = c.BindJSON(&department)
	if err != nil {
		panic(err)
	}

	//config.DB.First(&department, "id = ?", id)
	result := db.Model(models.Department{}).Debug().Where("id = ?", id).Updates(&department)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    department,
	})
}

func DeleteDepartment(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}

	stringId := c.Param("id")

	id, _ := strconv.Atoi(stringId)

	var department models.Department

	resDB := db.Delete(&department, "id = ?", id)

	if resDB.RowsAffected == 0 {
		c.JSON(http.StatusOK, gin.H{
			"MSG": "Data tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": http.StatusOK,
		"data":    "Berhasil Delete Data",
	})
}
