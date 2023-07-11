package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     Archive
	//Employee    []EmployeeInventory
}

type InventoryResponse struct {
	ID          uint
	Name        string `json:"name"`
	Description string `json:"description"`
	Archive     ArchiveResponse
	//Employee    []EmployeeInventory
}
