package controller

import (
	"ning/go-dashboard/features/bar_chart/entities"
	"ning/go-dashboard/features/bar_chart/service"
	"ning/go-dashboard/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type ChartExpenseController struct {
	service *service.ChartExpenseService
}

func NewChartExpenseController(service *service.ChartExpenseService) *ChartExpenseController {
	return &ChartExpenseController{service: service}
}

func (c *ChartExpenseController) GetExpenseData(ctx *fiber.Ctx) error {
	var body entities.ChartCilent

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.ChartExpenseResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []entities.DiseaseExpenseData{},
			})
	}

	dateStart, err := utils.ParseDate(body.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.ChartExpenseResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []entities.DiseaseExpenseData{},
			})
	}

	dateEnd, err := utils.ParseDate(body.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.ChartExpenseResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []entities.DiseaseExpenseData{},
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

	body.StartDate = dateStart.Format("2006-01-02")
	body.EndDate = dateEnd.Format("2006-01-02")

	expenseData, err := c.service.GetExpenseData(chartRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&entities.ChartExpenseResponse{
				Status:  false,
				Message: "Error for get expense data",
				Result:  []entities.DiseaseExpenseData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&entities.ChartExpenseResponse{
			Status:  true,
			Message: "Success",
			Result:  expenseData,
		})
}
