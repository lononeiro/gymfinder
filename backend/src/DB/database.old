package DB

import (
	"log"

	"github.com/lononeiro/gymfinder/backend/src/config"
	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func ConnectAndMigrate() {
	dsn := config.LoadDSN()

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	// AutoMigrate faz as migrations dos models
	err = database.AutoMigrate(
		&model.Academia{},
		&model.Usuario{},
		&model.Comentario{},
		&model.Imagem{},
	)
	if err != nil {
		log.Fatal("migration failed:", err)
	}

	DataBase = database

	log.Println("Database connection successful and migrations ran!")
}
