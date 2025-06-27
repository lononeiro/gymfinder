package DB

import (
	"log"

	"github.com/lononeiro/gymfinder/backend/src/config"
	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectAndMigrate() {
	dsn := config.LoadDSN()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// AutoMigrate faz as migrations dos models
	err = database.AutoMigrate(
		&model.Academia{},
	)
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	DB = database

	log.Println("Database connection successful and migrations ran!")
}
