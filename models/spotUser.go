package models

import "time"

type SpotUser struct {
	ID            uint `gorm:"primaryKey"`
	SpotID        uint
	UserID        uint
	InTime        time.Time
	OutTime       time.Time
	Active        bool
	AllocatedDate time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Spot          Spot `gorm:"foreignKey:SpotID"`
	User          User `gorm:"foreignKey:UserID"`
}
