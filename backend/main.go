package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/router"
	"github.com/rs/cors"
)

func main() {
	DB.ConnectAndMigrate()
	fmt.Println("Aplicação começou")

	r := router.InitializeRoutes()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("Servidor rodando em http://localhost:8081")
	err := http.ListenAndServe(":8081", handler)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}