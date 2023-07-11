package routes

import (
	"github.com/gin-gonic/gin"
	"golang_basic_gin/config"
	"golang_basic_gin/models"
	"gorm.io/gorm/clause"
	"net/http"
)

func GetInventory(c *gin.Context) {
	db, err := config.InitDB()
	if err != nil {
		panic(err.Error())
	}
	inventory := []models.Inventory{}
	db.Preload(clause.Associations).Find(&inventory)

	inventoryResponse := []models.InventoryResponse{}
	//archiveResponse := models.ArchiveResponse{}

	for _, p := range inventory {
		employee := models.InventoryResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Archive: models.ArchiveResponse{
				ID:          p.Archive.ID,
				Name:        p.Archive.Name,
				Description: p.Archive.Description,
				InventoryID: p.Archive.InventoryID,
			},
		}

		inventoryResponse = append(inventoryResponse, employee)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Get Position",
		"data":    inventoryResponse,
	})
}
