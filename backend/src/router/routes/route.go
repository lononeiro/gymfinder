package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI      string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := append(AcademiaRoutes, UsuarioRoutes...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Function).Methods(rota.Method)
	}
	return r
}
