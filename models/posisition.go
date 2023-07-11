package models

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name         string `gorm:"column:name" json:"name"`
	Code         string `gorm:"column:code" json:"code"`
	DepartmentId uint   `gorm:"column:department_id" json:"department_id"`
	//Employees    []Employee `json:"employees"`
}

type PositionRequest struct {
	Name         string `json:"name"`
	Code         string `json:"code"`
	DepartmentId uint   `json:"department_id"`
}

type PositionResponse struct {
	ID           uint               `json:"id"`
	Name         string             `json:"name"`
	Code         string             `json:"code"`
	DepartmentId uint               `json:"department_id"`
	Departments  DepartmentResponse `json:"department"`
}

type PositionResponseInnerJoin struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
