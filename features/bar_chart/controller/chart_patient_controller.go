package controller

import (
	"ning/go-dashboard/features/bar_chart/entities"
	"ning/go-dashboard/features/bar_chart/service"
	"ning/go-dashboard/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type ChartPatientController struct {
	service *service.ChartPatientService
}

func NewChartPatientController(service *service.ChartPatientService) *ChartPatientController {
	return &ChartPatientController{service: service}
}

func (c *ChartPatientController) GetPatientData(ctx *fiber.Ctx) error {
	var body entities.ChartCilent

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.ChartPatientResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []entities.DiseasePatientData{},
			})
	}

	dateStart, err := utils.ParseDate(body.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.ChartPatientResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []entities.DiseasePatientData{},
			})
	}

	dateEnd, err := utils.ParseDate(body.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.ChartPatientResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []entities.DiseasePatientData{},
			})
	}

	chartRequest := entities.ChartRequest{
		StartDate: dateStart,
		EndDate:   dateEnd,
		Area:      body.Area,
		Province:  body.Province,
		District:  body.District,
		Hcode:     body.Hcode,
	}

	patientData, err := c.service.GetPatientData(chartRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&entities.ChartPatientResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []entities.DiseasePatientData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&entities.ChartPatientResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
