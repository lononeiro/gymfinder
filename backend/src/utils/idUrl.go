package utils

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RetornarIdURL(w http.ResponseWriter, r *http.Request) (uint, bool) {

	vars := mux.Vars(r)

	idStr := vars["id"]

	if idStr == "" {
		respondWithJSONError(w, "ID não fornecido na URL", http.StatusBadRequest)
		return 0, false
	}
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		respondWithJSONError(w, "ID inválido: deve ser um número positivo", http.StatusBadRequest)
		return 0, false
	}

	return uint(id), true

}

func respondWithJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
