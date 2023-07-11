package db

import (
	"fmt"
	"github.com/ekantbajaj/duty-allocation/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// InitDB initializes the database connection
func Init() error {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.name"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
	)
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// AutoMigrate models if needed
	// db.AutoMigrate(&models.Spot{}, &models.User{}, &models.SpotUser{})
	return runMigrations()
}

func GetDB() *gorm.DB {
	return db
}

// Close closes the database connection.
func Close() {
	if db != nil {
		dbSQL, err := db.DB()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
			return
		}
		err = dbSQL.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
			return
		}
	}
}

// runMigrations runs the database migrations.
func runMigrations() error {
	err := db.AutoMigrate(&models.Spot{}, &models.User{}, &models.SpotUser{}, &models.Role{})
	if err != nil {
		return err
	}

	// Add more AutoMigrate statements for other models as needed

	return nil
}
