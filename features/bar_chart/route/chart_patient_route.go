package route

import (
	"ning/go-dashboard/features/bar_chart/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App, controller *controller.ChartPatientController) {
	app.Post("/bar-chart/pt-data", controller.GetPatientData)
}
