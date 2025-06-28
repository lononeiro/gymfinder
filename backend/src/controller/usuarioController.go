package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"github.com/lononeiro/gymfinder/backend/src/repository"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		fmt.Println("Erro ao decodificar o corpo da requisição:", err)
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Adicione verificação de erro
	err = repository.CreateUsuario(usuario)
	if err != nil {
		fmt.Println("Erro ao criar usuário:", err)
		http.Error(w, "Erro ao criar usuário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Usuário adicionado: %+v\n", usuario)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		fmt.Println("Erro ao decodificar o corpo da requisição:", err)
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	usuario, err = repository.LoginUsuario(usuario.Email, usuario.Senha)
	if err != nil {
		fmt.Println("Erro ao fazer login:", err)
		http.Error(w, "Erro ao fazer login: "+err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(uint(usuario.ID), usuario.Admin)
	if err != nil {
		fmt.Println("Erro ao gerar token:", err)
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Login realizado: %s (ID: %d)\n", usuario.Nome, usuario.ID)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":    token,
		"nome":     usuario.Nome,
		"id":       usuario.ID,
		"is_admin": usuario.Admin,
	})
}

func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Obter usuários do repositório
	usuarios := repository.ListarUsuarios()

	// Configurar cabeçalhos e status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Criar estrutura de resposta
	response := map[string]interface{}{
		"usuarios": usuarios,
		"count":    len(usuarios),
	}

	// Codificar resposta com formatação (indentação)
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("[ERRO] Falha ao codificar resposta JSON: %v", err)
		http.Error(w, "Erro ao formatar resposta", http.StatusInternalServerError)
		return
	}

	// Log de sucesso
	log.Printf("[INFO] Listagem realizada - %d usuários retornados", len(usuarios))

	// Escrever resposta
	w.Write(jsonData)
}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	// Implementação futura
}

func ApagarUsuario(w http.ResponseWriter, r *http.Request) {
	id, _ := utils.RetornarIdURL(w, r)
	if id == 0 {
		fmt.Println("ID inválido para apagar usuário")
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err := repository.ApagarUsuario(id)
	if err != nil {
		fmt.Println("Erro ao apagar usuario:", err)
		http.Error(w, "Erro ao apagar usuario: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Usuário apagado: ID %d\n", id)
	w.WriteHeader(http.StatusNoContent)
}
