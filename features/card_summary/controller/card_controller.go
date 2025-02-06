package controller

import (
	"ning/go-dashboard/features/card_summary/entities/dto"
	"ning/go-dashboard/features/card_summary/service"
	"ning/go-dashboard/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type CardCtrl struct {
	service *service.CardService
}

func NewCardCtrl(service *service.CardService) *CardCtrl {
	return &CardCtrl{service: service}
}

func (c *CardCtrl) GetCardData(ctx *fiber.Ctx) error {
	var body dto.CardCilent

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.CardResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []dto.CardData{},
			})
	}

	dateStart, err := utils.ParseDate(body.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.CardResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []dto.CardData{},
			})
	}

	dateEnd, err := utils.ParseDate(body.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&dto.CardResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []dto.CardData{},
			})
	}

	cardRequest := dto.CardRequest{
		StartDate: dateStart,
		EndDate:   dateEnd,
		Area:      body.Area,
		Province:  body.Province,
		District:  body.District,
		Hcode:     body.Hcode,
	}

	body.StartDate = dateStart.Format("2006-01-02")
	body.EndDate = dateEnd.Format("2006-01-02")

	cardData, err := c.service.GetCardData(cardRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			&dto.CardResponse{
				Status:  false,
				Message: "Error for get card data",
				Result:  []dto.CardData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&dto.CardResponse{
			Status:  true,
			Message: "Success",
			Result:  cardData,
		})
}
