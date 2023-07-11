package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	Name       string                    `json:"name"`
	Address    string                    `json:"address"`
	Email      string                    `json:"email"`
	PositionId uint                      `json:"position_id"`
	Position   PositionResponseInnerJoin `json:"position"`
}

type EmployeeRequest struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Email      string `json:"email"`
	PositionId uint   `json:"position_id"`
}

type EmployeeResponse struct {
	ID         uint                      `json:"id"`
	Name       string                    `json:"name"`
	Address    string                    `json:"address"`
	Email      string                    `json:"email"`
	PositionId uint                      `json:"department_id"`
	Position   PositionResponseInnerJoin `json:"position"`
}
