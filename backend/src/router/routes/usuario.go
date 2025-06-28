package routes

import (
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/controller"
	"github.com/lononeiro/gymfinder/backend/src/utils/middleware"
)

var UsuarioRoutes = []Route{
	{
		URI:    "/usuario",
		Method: "GET",
		Function: func(w http.ResponseWriter, r *http.Request) {
			middleware.AdminOnly(http.HandlerFunc(controller.ListarUsuarios)).ServeHTTP(w, r)
		},
	},
	{
		URI:    "/usuario",
		Method: "DELETE",
		Function: func(w http.ResponseWriter, r *http.Request) {
			middleware.AdminOnly(http.HandlerFunc(controller.ApagarUsuario)).ServeHTTP(w, r)
		},
	},

	{
		URI:      "/usuario",
		Method:   "POST",
		Function: controller.AdicionarUsuario,
	},
	{
		URI:      "/usuario/login",
		Method:   "POST",
		Function: controller.LoginHandler,
	},
}
