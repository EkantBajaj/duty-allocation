package models

import (
	"time"
)

type User struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `json:"name" binding:"required"`
	RelativeName    string `json:"relative_name" binding:"required"`
	Gender          string `json:"gender" binding:"required"`
	MobileNumber    string `json:"mobile_number" binding: "required" gorm:"unique"`
	BadgeID         string `json:"badge_id" gorm:"unique"`
	Initiated       bool   `json:"initiated" default:"false"`
	Address         string `json:"address" binding:"required"`
	Email           string `json:"email" gorm:"unique"`
	City            string `json:"city" binding:"required"`
	PinCode         string `json:"pin_code" binding:"required"`
	EmergencyNumber string `json:"emergency_number" binding:"required"`
	BirthDate       string `json:"birth_date"`
	InitiationDate  string `json:"initiation_date"`
	Qualification   string `json:"qualification" binding:"required"`
	Profession      string `json:"profession" binding:"required"`
	MaritalStatus   string `json:"marital_status" binding:"required"`
	BloodGroup      string `json:"blood_group" binding:"required"`
	Department      string `json:"department" binding:"required"`
	AadharNumber    string `json:"aadhar_number" binding:"required"`
	ZoneBadgeID     string `json:"zone_badge_id"`
	ZoneDepartment  string `json:"zone_department"`
	IntroducedBy    string `json:"introduced_by" binding:"required"`
	Center          string `json:"center" binding:"required"`
	SubCenter       string `json:"sub_center" binding:"required"`
	PhotoBadge      bool   `json:"photo_badge" default:"false"`
	BadgePrinted    bool   `json:"badge_printed" default:"false"`
	Remarks         string `json:"remarks" binding:"required"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Roles           []Role `gorm:"many2many:user_roles;"`
	Password        string `json:"password" binding:"required"`
}

type ActiveUsersBySpotID struct {
	ID              uint
	Name            string
	Gender          string
	BadgeID         string `gorm:"unique"`
	MobileNumber    string `gorm:"unique"`
	EmergencyNumber string
	Initiated       bool
	InTime          time.Time
	OutTime         time.Time
	InTimeString    string // Formatted time string for InTime
	OutTimeString   string
}
type Login struct {
	BadgeID  string `json:"badge_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// create user request with fields name email password badgeid and if badgeid is not given then make it last four digit of adhar
