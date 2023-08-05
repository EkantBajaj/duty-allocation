package models

import (
	"time"
)

type User struct {
	ID              uint `gorm:"primaryKey"`
	Name            string
	RelativeName    string
	Gender          string
	MobileNumber    string `gorm:"unique"`
	BadgeID         string `gorm:"unique"`
	Initiated       bool
	Address         string
	Email           string
	City            string
	PinCode         string
	EmergencyNumber string
	BirthDate       time.Time
	InitiationDate  time.Time
	Qualification   string
	Profession      string
	MaritalStatus   string
	BloodGroup      string
	Department      string
	AadharNumber    string
	ZoneBadgeID     string
	ZoneDepartment  string
	IntroducedBy    string
	Center          string
	SubCenter       string
	PhotoBadge      bool
	BadgePrinted    bool
	Remarks         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Roles           []Role `gorm:"many2many:user_roles;"`
	Password        string `json:"password"`
}

type ActiveUsersBySpotID struct {
	ID        uint
	Name      string
	Gender    string
	BadgeID   string `gorm:"unique"`
	Initiated bool
}
