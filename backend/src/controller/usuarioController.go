package controller

import (
	"encoding/json"
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"github.com/lononeiro/gymfinder/backend/src/repository"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

func AdicionarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}
	repository.CreateUsuario(usuario)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	usuario, err = repository.LoginUsuario(usuario.Email, usuario.Senha)
	if err != nil {
		http.Error(w, "Erro ao fazer login: "+err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(uint(usuario.ID), usuario.Admin)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"token":    token,
		"nome":     usuario.Nome,
		"id":       usuario.ID,
		"is_admin": usuario.Admin,
	})
}
