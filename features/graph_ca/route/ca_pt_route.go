package route

import (
	"ning/go-dashboard/features/graph_ca/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App, controller *controller.GraphCaPtCtrl) {
	app.Post("/graph-ca-pt", controller.GetGraphCaPtData)
}
