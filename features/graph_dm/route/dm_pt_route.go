package route

import (
	"ning/go-dashboard/features/graph_dm/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App, controller *controller.GraphDmPtCtrl) {
	app.Post("/graph-dm-pt", controller.GetGraphDmPtData)
}
