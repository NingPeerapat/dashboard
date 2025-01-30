package service

import (
	"fmt"
	"ning/go-dashboard/features/card_summary/entities"
	"ning/go-dashboard/features/card_summary/repository"
	"ning/go-dashboard/pkg/utils"
)

type CardService struct {
	repo *repository.CardRepository
}

func NewCardService(repo *repository.CardRepository) *CardService {
	return &CardService{repo: repo}
}

func (service *CardService) GetCardData(body entities.CardRequest) ([]entities.CardData, error) {

	cardData, err := service.repo.GetAllData(entities.CardRequest{
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
		Area:      body.Area,
		Province:  body.Province,
		District:  body.District,
		Hcode:     body.Hcode,
	})

	if err != nil {
		return nil, fmt.Errorf("error fetching card data: %v", err)
	}

	cidCountData, err := service.repo.GetCidCountData(entities.CardRequest{
		StartDate: body.StartDate,
		EndDate:   body.EndDate,
		Area:      body.Area,
		Province:  body.Province,
		District:  body.District,
		Hcode:     body.Hcode,
	})
	if err != nil {
		return nil, fmt.Errorf("error fetching CID count data: %v", err)
	}

	finalResult := []entities.CardData{
		calculateCardDataData(*cardData, *cidCountData),
	}

	return finalResult, nil
}

func calculateCardDataData(cardRawData entities.CardRawData, cidCountData entities.CidCountData) entities.CardData {
	var avgService float64
	var avgExpense float64
	expense := float64(cardRawData.DmExpense + cardRawData.HtExpense + cardRawData.CaExpense + cardRawData.CopdExpense + cardRawData.PsyExpense + cardRawData.HdCvdExpense)
	if cidCountData.CidCount > 0 {
		avgService = float64(cardRawData.ServiceCount) / float64(cidCountData.CidCount)
		avgExpense = expense / float64(cidCountData.CidCount)
	}

	return entities.CardData{
		ServiceCount: float64(cardRawData.ServiceCount),
		PatientCount: float64(cidCountData.CidCount),
		Expense:      utils.RoundToTwoDecimalPlaces(expense),
		HcodeCount:   cardRawData.HcodeCount,
		AvgService:   utils.RoundToTwoDecimalPlaces(avgService),
		AvgExpense:   utils.RoundToTwoDecimalPlaces(avgExpense),
	}
}
