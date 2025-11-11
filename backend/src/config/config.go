package config

import (
	"fmt"
	// "log"
	"os"

	"github.com/joho/godotenv"
	// "gorm.io/driver/postgres"

)

func LoadDSN() string {
	_ = godotenv.Load()

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		host, user, pass, dbname, port,
	)

	return dsn
}
