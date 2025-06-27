package routes

import (
	"net/http"

	"github.com/lononeiro/gymfinder/backend/src/controller"
	"github.com/lononeiro/gymfinder/backend/src/utils/middleware"
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
		URI:    "/academia",
		Method: "DELETE",
		Function: func(w http.ResponseWriter, r *http.Request) {
			middleware.AdminOnly(http.HandlerFunc(controller.ApagarAcademia)).ServeHTTP(w, r)
		},
	},
}
