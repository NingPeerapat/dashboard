package controller

import (
	"ning/go-dashboard/features/card_summary/entities"
	"ning/go-dashboard/features/card_summary/service"
	"ning/go-dashboard/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type CardController struct {
	service *service.CardService
}

func NewCardController(service *service.CardService) *CardController {
	return &CardController{service: service}
}

func (c *CardController) GetCardData(ctx *fiber.Ctx) error {
	var body entities.CardCilent

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.CardResponse{
				Status:  false,
				Message: "Error in parsing request body",
				Result:  []entities.CardData{},
			})
	}

	dateStart, err := utils.ParseDate(body.StartDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.CardResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []entities.CardData{},
			})
	}

	dateEnd, err := utils.ParseDate(body.EndDate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			&entities.CardResponse{
				Status:  false,
				Message: err.Error(),
				Result:  []entities.CardData{},
			})
	}

	cardRequest := entities.CardRequest{
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
			&entities.CardResponse{
				Status:  false,
				Message: "Error for get card data",
				Result:  []entities.CardData{},
			})
	}

	return ctx.Status(fiber.StatusOK).JSON(
		&entities.CardResponse{
			Status:  true,
			Message: "Success",
			Result:  cardData,
		})
}
