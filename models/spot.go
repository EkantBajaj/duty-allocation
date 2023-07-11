package models

import "time"

type Spot struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	TotalPeople int
	MinPeople   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ActiveSpotUserCount struct {
	ID          uint   // Spot ID
	Name        string // Spot name
	UserCount   int    // Count of users in the spot
	TotalPeople int
	MinPeople   int
}
