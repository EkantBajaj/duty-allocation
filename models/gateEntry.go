package models

import "time"

type GateEntry struct {
	ID        uint   `gorm:"primaryKey"`
	GateID    uint   `json:"gate_id" binding:"required"`
	Count     int    `json:"count" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}