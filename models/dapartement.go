package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name       string     `json:"name"`
	Code       string     `json:"code"`
	Posistions []Position `json:"posistions"`
}

type DepartmentRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type DepartmentResponse struct {
	ID         uint                        `json:"id"`
	Name       string                      `json:"name"`
	Code       string                      `json:"code"`
	Posistions []PositionResponseInnerJoin `json:"positions"`
}
