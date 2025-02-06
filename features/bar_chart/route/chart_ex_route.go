package route

import (
	"ning/go-dashboard/features/bar_chart/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterExpenseRoutes(app *fiber.App, controller *controller.ChartExCtrl) {
	app.Post("/chart-ex", controller.GetChartExData)
}
