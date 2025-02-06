package controller

import (
	"ning/go-dashboard/features/graph_dm/entities/dto"
	"ning/go-dashboard/features/graph_dm/service"

	"github.com/gofiber/fiber/v2"
)

type GraphDmExCtrl struct {
	service *service.GraphDmExService
}

func NewGraphDmExCtrl(service *service.GraphDmExService) *GraphDmExCtrl {
	return &GraphDmExCtrl{service: service}
}

func (c *GraphDmExCtrl) GetGraphDmExData(ctx *fiber.Ctx) error {
	var body dto.DmRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.DmResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []dto.DmData{},
			})
	}

	CaRequest := dto.DmRequest{
		Year:     2024,
		Area:     body.Area,
		Province: body.Province,
		District: body.District,
		Hcode:    body.Hcode,
	}

	expenseData, err := c.service.GetGraphDmExData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.DmResponse{
				Status:  false,
				Message: "Error for get expense data",
				Result:  []dto.DmData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.DmResponse{
			Status:  true,
			Message: "Success",
			Result:  expenseData,
		})
}
