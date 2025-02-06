package route

import (
	"ning/go-dashboard/features/card_summary/controller"

	"github.com/gofiber/fiber/v2"
)

func RegisterCardRoutes(app *fiber.App, controller *controller.CardCtrl) {
	app.Post("/card", controller.GetCardData)
}
