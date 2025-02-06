package controller

import (
	"ning/go-dashboard/features/graph_ca/entities/dto"
	"ning/go-dashboard/features/graph_ca/service"

	"github.com/gofiber/fiber/v2"
)

type GraphCaPtCtrl struct {
	service *service.GraphCaPtService
}

func NewGraphCaPtCtrl(service *service.GraphCaPtService) *GraphCaPtCtrl {
	return &GraphCaPtCtrl{service: service}
}

func (c *GraphCaPtCtrl) GetGraphCaPtData(ctx *fiber.Ctx) error {
	var body dto.CaRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.CaResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []dto.CaData{},
			})
	}

	CaRequest := dto.CaRequest{
		Year:     2024,
		Area:     body.Area,
		Province: body.Province,
		District: body.District,
		Hcode:    body.Hcode,
	}

	patientData, err := c.service.GetGraphCaPtData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.CaResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []dto.CaData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.CaResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
