package models

import "time"

type User struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Gender       string
	MobileNumber string `gorm:"unique"`
	BadgeID      string `gorm:"unique"`
	Initiated    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Roles        []Role `gorm:"many2many:user_roles;"`
	Password     string `json:"password"`
}

type ActiveUsersBySpotID struct {
	ID        uint
	Name      string
	Gender    string
	BadgeID   string `gorm:"unique"`
	Initiated bool
}
