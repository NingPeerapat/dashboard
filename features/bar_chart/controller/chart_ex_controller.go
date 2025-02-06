package controller

import (
	"ning/go-dashboard/features/bar_chart/entities/dto"
	"ning/go-dashboard/features/bar_chart/service"
	"ning/go-dashboard/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type ChartExCtrl struct {
	service *service.ChartExService
}

func NewChartExCtrl(service *service.ChartExService) *ChartExCtrl {
	return &ChartExCtrl{service: service}
}

func (c *ChartExCtrl) GetChartExData(ctx *fiber.Ctx) error {
	var body dto.ChartCilent

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.ChartExpenseResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []dto.ChartExpenseData{},
			})
	}

	dateStart, err := utils.ParseDate(body.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.ChartExpenseResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []dto.ChartExpenseData{},
			})
	}

	dateEnd, err := utils.ParseDate(body.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.ChartExpenseResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []dto.ChartExpenseData{},
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

	body.StartDate = dateStart.Format("2006-01-02")
	body.EndDate = dateEnd.Format("2006-01-02")

	expenseData, err := c.service.GetChartExData(chartRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.ChartExpenseResponse{
				Status:  false,
				Message: "Error for get expense data",
				Result:  []dto.ChartExpenseData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.ChartExpenseResponse{
			Status:  true,
			Message: "Success",
			Result:  expenseData,
		})
}
