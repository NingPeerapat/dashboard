package controller

import (
	"ning/go-dashboard/features/graph_disease/entities"
	"ning/go-dashboard/features/graph_disease/service"

	"github.com/gofiber/fiber/v2"
)

type DiseasePatientController struct {
	service *service.DiseasePatientService
}

func NewDiseasePatientController(service *service.DiseasePatientService) *DiseasePatientController {
	return &DiseasePatientController{service: service}
}

func (c *DiseasePatientController) GetPatientData(ctx *fiber.Ctx) error {
	var body entities.DiseaseRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.DiseaseResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []entities.DiseaseData{},
			})
	}

	CaRequest := entities.DiseaseRequest{
		Year:     2024,
		Area:     body.Area,
		Province: body.Province,
		District: body.District,
		Hcode:    body.Hcode,
	}

	patientData, err := c.service.GetPatientData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&entities.DiseaseResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []entities.DiseaseData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&entities.DiseaseResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
