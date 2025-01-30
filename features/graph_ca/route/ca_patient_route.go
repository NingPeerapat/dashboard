package route

import (
	"ning/go-dashboard/features/graph_ca/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App, controller *controller.GraphCaPatientController) {
	app.Post("/graph/ca-pt", controller.GetPatientData)
}
