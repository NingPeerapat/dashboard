package controller

import (
	"ning/go-dashboard/features/graph_disease/entities/dto"
	"ning/go-dashboard/features/graph_disease/service"

	"github.com/gofiber/fiber/v2"
)

type GraphDiseaseExCtrl struct {
	service *service.GraphDiseaseExService
}

func NewGraphDiseaseExCtrl(service *service.GraphDiseaseExService) *GraphDiseaseExCtrl {
	return &GraphDiseaseExCtrl{service: service}
}

func (c *GraphDiseaseExCtrl) GetGraphDiseaseExData(ctx *fiber.Ctx) error {
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

	expenseData, err := c.service.GetGraphDiseaseExData(CaRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.DiseaseResponse{
				Status:  false,
				Message: "Error for get expense data",
				Result:  []dto.DiseaseData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.DiseaseResponse{
			Status:  true,
			Message: "Success",
			Result:  expenseData,
		})
}
