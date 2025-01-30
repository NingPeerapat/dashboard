package route

import (
	"ning/go-dashboard/features/graph_ca/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterExpenseRoutes(app *fiber.App, controller *controller.GraphCaExpenseController) {
	app.Post("/graph/ca-ex", controller.GetExpenseData)
}
