package controller

import (
	"ning/go-dashboard/features/graph_ca/entities"
	"ning/go-dashboard/features/graph_ca/service"

	"github.com/gofiber/fiber/v2"
)

type GraphCaPatientController struct {
	service *service.GraphCaPatientService
}

func NewGraphCaPatientController(service *service.GraphCaPatientService) *GraphCaPatientController {
	return &GraphCaPatientController{service: service}
}

func (c *GraphCaPatientController) GetPatientData(ctx *fiber.Ctx) error {
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

	patientData, err := c.service.GetPatientData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&entities.CaResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []entities.CaData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&entities.CaResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
