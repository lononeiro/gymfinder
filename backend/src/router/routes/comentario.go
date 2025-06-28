package routes

import (
	"github.com/lononeiro/gymfinder/backend/src/controller"
	"github.com/lononeiro/gymfinder/backend/src/utils"
)

var ComentarioRoutes = []Route{
	{
		URI:      "/academia/{id}/comentario",
		Method:   "GET",
		Function: controller.ListarComentarios,
	},
	{
		URI:      "/academia/{id}/comentario",
		Method:   "POST",
		Function: applyMiddlewares(controller.CriarComentario, utils.AuthMiddleware),
	},
	{
		URI:      "/comentario/{id}",
		Method:   "DELETE",
		Function: applyMiddlewares(controller.ApagarComentario, utils.AuthMiddleware),
	},
	{
		URI:      "/comentario/{id}",
		Method:   "PUT",
		Function: applyMiddlewares(controller.EditarComentario, utils.AuthMiddleware),
	},
}
