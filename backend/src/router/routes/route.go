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
	rotas := append(append(AcademiaRoutes, UsuarioRoutes...), ComentarioRoutes...)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Function).Methods(rota.Method)
	}
	return r
}

func applyMiddlewares(handler http.HandlerFunc, middlewares ...func(http.Handler) http.Handler) http.HandlerFunc {
	h := http.Handler(handler)
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h.ServeHTTP
}
