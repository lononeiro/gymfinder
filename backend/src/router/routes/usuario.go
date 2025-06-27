package routes

import (
	"github.com/lononeiro/gymfinder/backend/src/controller"
)

var UsuarioRoutes = []Route{
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
