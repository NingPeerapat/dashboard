package route

import (
	"ning/go-dashboard/features/graph_disease/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterExpenseRoutes(app *fiber.App, controller *controller.DiseaseExpenseController) {
	app.Post("/graph/disease-ex", controller.GetExpenseData)
}
