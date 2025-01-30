package controller

import (
	"ning/go-dashboard/features/graph_ca/entities"
	"ning/go-dashboard/features/graph_ca/service"

	"github.com/gofiber/fiber/v2"
)

type GraphCaExpenseController struct {
	service *service.GraphCaExpenseService
}

func NewGraphCaExpenseController(service *service.GraphCaExpenseService) *GraphCaExpenseController {
	return &GraphCaExpenseController{service: service}
}

func (c *GraphCaExpenseController) GetExpenseData(ctx *fiber.Ctx) error {
	var body entities.CaRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.CaResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []entities.CaData{},
			})
	}

	CaRequest := entities.CaRequest{
		Year:     2024,
		Area:     body.Area,
		Province: body.Province,
		District: body.District,
		Hcode:    body.Hcode,
	}

	expenseData, err := c.service.GetExpenseData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&entities.CaResponse{
				Status:  false,
				Message: "Error for get expense data",
				Result:  []entities.CaData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&entities.CaResponse{
			Status:  true,
			Message: "Success",
			Result:  expenseData,
		})
}
