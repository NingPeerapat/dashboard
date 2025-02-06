package route

import (
	"ning/go-dashboard/features/graph_ca/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterExpenseRoutes(app *fiber.App, controller *controller.GraphCaExCtrl) {
	app.Post("/graph-ca-ex", controller.GetGraphCaExData)
}
