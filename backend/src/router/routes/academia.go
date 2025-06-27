package routes

import (
	"github.com/lononeiro/gymfinder/backend/src/controller"
)

var AcademiaRoutes = []Route{
	{
		URI:      "/academia",
		Method:   "POST",
		Function: controller.AdicionarAcademia,
	},
	{
		URI:      "/academias",
		Method:   "GET",
		Function: controller.ListarAcademias,
	},

	{
		URI:      "/academia",
		Method:   "PUT",
		Function: controller.EditarAcademias,
	},
	{
		URI:      "/academia",
		Method:   "DELETE",
		Function: controller.ApagarAcademia,
	},
}
