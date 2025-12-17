package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"github.com/lononeiro/gymfinder/backend/src/repository"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

func CriarComentario(w http.ResponseWriter, r *http.Request) {
	var comentario model.Comentario

	idAcademia, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	err := json.NewDecoder(r.Body).Decode(&comentario)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	comentario.AcademiaID = idAcademia

	comentario.UsuarioID, err = utils.ExtractUserIDFromToken(r)

	fmt.Println(comentario.UsuarioID, comentario.AcademiaID)

	if err != nil {
		http.Error(w, "Erro ao extrair ID do usuário do token: "+err.Error(), http.StatusUnauthorized)
		return
	}

	err = repository.CriarComentario(comentario)
	if err != nil {
		http.Error(w, "Erro ao criar comentário: "+err.Error(), http.StatusInternalServerError)
		return
	}

}

func ApagarComentario(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.RetornarIdURL(w, r)

	if !ok {
		return
	}

	err := repository.ApagarComentario(id)
	if err != nil {
		http.Error(w, "Erro ao apagar comentário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func EditarComentario(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	var comentario model.Comentario
	err := json.NewDecoder(r.Body).Decode(&comentario)
	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = repository.EditarComentario(id, comentario)
	if err != nil {
		http.Error(w, "Erro ao editar comentário: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comentario)
}

func ListarComentarios(w http.ResponseWriter, r *http.Request) {
	academiaID, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	comentarios, err := repository.ListarComentariosPost(academiaID)
	if err != nil {
		http.Error(w, "Erro ao listar comentários: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comentarios)
}

func SelecionarUsuarioComentario(w http.ResponseWriter, r *http.Request) {
	id, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}
	usuario, err := repository.SelecionarUsuarioComentario(id)
	if err != nil {
		http.Error(w, "Erro ao selecionar usuário do comentário: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usuario)
}
