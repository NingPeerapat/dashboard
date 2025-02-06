package route

import (
	"ning/go-dashboard/features/graph_dm/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterExpenseRoutes(app *fiber.App, controller *controller.GraphDmExCtrl) {
	app.Post("/graph-dm-ex", controller.GetGraphDmExData)
}
