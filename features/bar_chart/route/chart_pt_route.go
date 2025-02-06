package route

import (
	"ning/go-dashboard/features/bar_chart/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterPatientRoutes(app *fiber.App, controller *controller.ChartPtCtrl) {
	app.Post("/chart-pt", controller.GetChartPtData)
}
