package DB

import (
	"log"

	"github.com/lononeiro/gymfinder/backend/src/config"
	"github.com/lononeiro/gymfinder/backend/src/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func ConnectAndMigrate() {
	dsn := config.LoadDSN()

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: false})

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
