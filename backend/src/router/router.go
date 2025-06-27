package router

import (
	"github.com/gorilla/mux"
	"github.com/lononeiro/gymfinder/backend/src/router/routes"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()

	return routes.Configurar(r)
}
