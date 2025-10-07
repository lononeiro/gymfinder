package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/model"
	"github.com/lononeiro/gymfinder/backend/src/repository"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

func AdicionarAcademia(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		model.Academia `json:"academia"`
		Imagem model.Imagem `json:"imagem"`
	}
	err := json.NewDecoder(r.Body).Decode(&payload)

	if err != nil {
		http.Error(w, "Erro ao decodificar o corpo da requisição: "+err.Error(), http.StatusBadRequest)
		return
	}
	academia, err := repository.CreateAcademia(payload.Academia, payload.Imagem)
	if err != nil {
		http.Error(w, "Erro ao criar academia: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Academia criada com sucesso",
		"academia":  academia,
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}


func ListarAcademias(w http.ResponseWriter, r *http.Request) {
	academias := repository.ListarAcademias()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]interface{}{
		"academias": academias,
	}
	jsonData, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(w, "Erro ao codificar a resposta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func EditarAcademias(w http.ResponseWriter, r *http.Request) {
	var academia model.Academia
	if err := json.NewDecoder(r.Body).Decode(&academia); err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Erro ao decodificar o corpo da requisição: %s"}`, err.Error()), http.StatusBadRequest)
		return
	}

	id, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	err := repository.EditarAcademias(id, academia)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Erro ao editar academia: %s"}`, err.Error()), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}
func ApagarAcademia(w http.ResponseWriter, r *http.Request) {

	id, ok := utils.RetornarIdURL(w, r)
	if !ok {
		return
	}

	err := repository.ApagarAcademia(id)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"Erro ao apagar academia: %s"}`, err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
