package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/router"
)

func main() {

	DB.ConnectAndMigrate()
	fmt.Println("Aplicação começou")

	r := router.InitializeRoutes()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
