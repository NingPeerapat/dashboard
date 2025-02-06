package controller

import (
	"ning/go-dashboard/features/graph_dm/entities/dto"
	"ning/go-dashboard/features/graph_dm/service"

	"github.com/gofiber/fiber/v2"
)

type GraphDmPtCtrl struct {
	service *service.GraphDmPtService
}

func NewGraphDmPtCtrl(service *service.GraphDmPtService) *GraphDmPtCtrl {
	return &GraphDmPtCtrl{service: service}
}

func (c *GraphDmPtCtrl) GetGraphDmPtData(ctx *fiber.Ctx) error {
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

	patientData, err := c.service.GetGraphDmPtData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.DmResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []dto.DmData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.DmResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
