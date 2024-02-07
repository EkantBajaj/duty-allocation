package models

type Gate struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `json:"name" binding:"required"`
}