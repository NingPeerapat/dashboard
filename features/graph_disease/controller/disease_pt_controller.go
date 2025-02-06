package controller

import (
	"ning/go-dashboard/features/graph_disease/entities/dto"
	"ning/go-dashboard/features/graph_disease/service"

	"github.com/gofiber/fiber/v2"
)

type GraphDiseasePtCtrl struct {
	service *service.GraphDiseasePtService
}

func NewGraphDiseasePtCtrl(service *service.GraphDiseasePtService) *GraphDiseasePtCtrl {
	return &GraphDiseasePtCtrl{service: service}
}

func (c *GraphDiseasePtCtrl) GetGraphDiseasePtData(ctx *fiber.Ctx) error {
	var body dto.DiseaseRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.DiseaseResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []dto.DiseaseData{},
			})
	}

	CaRequest := dto.DiseaseRequest{
		Year:     2024,
		Area:     body.Area,
		Province: body.Province,
		District: body.District,
		Hcode:    body.Hcode,
	}

	patientData, err := c.service.GetGraphDiseasePtData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.DiseaseResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []dto.DiseaseData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.DiseaseResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
