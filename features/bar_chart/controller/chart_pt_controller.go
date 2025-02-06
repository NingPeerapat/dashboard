package controller

import (
	"ning/go-dashboard/features/bar_chart/entities/dto"
	"ning/go-dashboard/features/bar_chart/service"
	"ning/go-dashboard/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type ChartPtCtrl struct {
	service *service.ChartPtService
}

func NewChartPtCtrl(service *service.ChartPtService) *ChartPtCtrl {
	return &ChartPtCtrl{service: service}
}

func (c *ChartPtCtrl) GetChartPtData(ctx *fiber.Ctx) error {
	var body dto.ChartCilent

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.ChartPatientResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []dto.ChartPatientData{},
			})
	}

	dateStart, err := utils.ParseDate(body.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.ChartPatientResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []dto.ChartPatientData{},
			})
	}

	dateEnd, err := utils.ParseDate(body.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.ChartPatientResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []dto.ChartPatientData{},
			})
	}

	chartRequest := dto.ChartRequest{
		StartDate: dateStart,
		EndDate:   dateEnd,
		Area:      body.Area,
		Province:  body.Province,
		District:  body.District,
		Hcode:     body.Hcode,
	}

	patientData, err := c.service.GetChartPtData(chartRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.ChartPatientResponse{
				Status:  false,
				Message: "Error for get patient data",
				Result:  []dto.ChartPatientData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.ChartPatientResponse{
			Status:  true,
			Message: "Success",
			Result:  patientData,
		})
}
