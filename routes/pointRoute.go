package routes

import (
	"api-geo-location/controllers"
	"api-geo-location/dtos"
	"api-geo-location/middlewares"

	"github.com/kataras/iris/v12"
)

type PointRoutes struct {
	pointController *controllers.PointController
}

func NewPointRoute() *PointRoutes {
	return &PointRoutes{
		pointController: controllers.NewPointController(),
	}
}

func (point *PointRoutes) Routes(app *iris.Application) iris.Party {
	routes := app.Party("/point")
	{
		routes.Post("/", middlewares.HandleValidatorError(&dtos.CreatePointDto{}), point.pointController.Create)
		routes.Get("/", point.pointController.GetAll)
		routes.Get("/{id:uint}", point.pointController.GetOne)
		routes.Put("/{id:uint}", middlewares.HandleValidatorError(&dtos.EditPointDto{}), point.pointController.Update)
		routes.Delete("/{id:uint}", point.pointController.Delete)
	}

	return routes
}
