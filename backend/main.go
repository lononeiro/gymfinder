package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/lononeiro/gymfinder/backend/src/DB"
	"github.com/lononeiro/gymfinder/backend/src/router"
	"github.com/lononeiro/gymfinder/backend/src/utils"
	"github.com/rs/cors"
)

func main() {

	// 🔥 1) CARREGAR VARIÁVEIS .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("⚠ Aviso: .env não encontrado ou não pôde ser carregado")
	} else {
		fmt.Println("✔ .env carregado com sucesso")
	}

	if err := utils.TestFilebaseConnection(); err != nil {
		fmt.Printf("Erro: %v\n", err)
	}

	// 2) conectar banco
	DB.ConnectAndMigrate()
	fmt.Println("Aplicação começou")

	// 3) rotas
	r := router.InitializeRoutes()

	// servir imagens locais (opcional se usar Filebase)
	r.PathPrefix("/uploads/").
		Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// 4) CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://gymfinder-1.onrender.com", "https://gymfinder-nine.vercel.app", "https://s3.filebase.com"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// 5) iniciar servidor
	fmt.Println("Servidor rodando em http://localhost:8081")
	err = http.ListenAndServe(":8081", handler)
	if err != nil {
		log.Fatal("Erro ao iniciar o servidor:", err)
	}
}
