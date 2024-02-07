package models

import "time"

type GateCount struct {
	ID        uint   `gorm:"primaryKey"`
	GateID    uint   `json:"gate_id" binding:"required"`
	Count     int   `json:"count" binding:"required"`
	Date      time.Time `json:"date"`
	Active   bool   `json:"active"`
}

type GateCountWithName struct {
	GateCount
	Name string `json:"name"`
}