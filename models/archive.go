package models

import (
	"gorm.io/gorm"
)

type Archive struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	InventoryID uint   `json:"inventory_id"`
}

type ArchiveResponse struct {
	ID          uint
	Name        string `json:"name"`
	Description string `json:"description"`
	InventoryID uint   `json:"inventory_id"`
}
