package routes

import (
	"github.com/lononeiro/gymfinder/backend/src/controller"
	"github.com/lononeiro/gymfinder/backend/src/utils"
	"github.com/lononeiro/gymfinder/backend/src/utils/middleware"
)

var AcademiaRoutes = []Route{
	{
		URI:      "/academia",
		Method:   "POST",
		Function: applyMiddlewares(controller.AdicionarAcademia),
	},
	{
		URI:      "/academias",
		Method:   "GET",
		Function: controller.ListarAcademias,
	},
	{
		URI:      "/academia/{id}",
		Method:   "PUT",
		Function: applyMiddlewares(controller.EditarAcademias, utils.AuthMiddleware, middleware.AdminOnly),
	},
	{
		URI:      "/academia/{id}",
		Method:   "DELETE",
		Function: applyMiddlewares(controller.ApagarAcademia, utils.AuthMiddleware, middleware.AdminOnly),
	},
}
