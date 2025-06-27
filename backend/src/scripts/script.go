package scripts

import (
	"fmt"
	"net/http"
	"strconv"
)

// Retorna o ID da URL ou responde erro JSON e retorna 0, false
func RetornarIdURL(w http.ResponseWriter, r *http.Request) (uint, bool) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, `{"error":"ID não informado"}`, http.StatusBadRequest)
		return 0, false
	}

	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"ID inválido: %s"}`, err.Error()), http.StatusBadRequest)
		return 0, false
	}
	return uint(idUint64), true
}
