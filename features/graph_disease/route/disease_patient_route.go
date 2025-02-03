package route

import (
	"ning/go-dashboard/features/graph_disease/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App, controller *controller.DiseasePatientController) {
	app.Post("/graph/disease-pt", controller.GetPatientData)
}
