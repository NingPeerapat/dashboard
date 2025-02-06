package service

import (
	"fmt"
	"ning/go-dashboard/features/card_summary/entities/dao"
	"ning/go-dashboard/features/card_summary/entities/dto"
	"ning/go-dashboard/features/card_summary/repository"
	"ning/go-dashboard/pkg/utils"
	"time"
)

type CardService struct {
	repo *repository.CardRepo
}

func NewCardService(repo *repository.CardRepo) *CardService {
	return &CardService{repo: repo}
}

func (service *CardService) GetCardData(body dto.CardRequest) ([]dto.CardData, error) {

	endDate := time.Now()
	formattedDate := endDate.Format("2006-01-02")
	startDateFormatted := body.StartDate.Format("2006-01-02")
	endDateFormatted := body.EndDate.Format("2006-01-02")

	if startDateFormatted == "2024-01-07" &&
		endDateFormatted == formattedDate &&
		body.Area == "" &&
		body.Province == "" &&
		body.District == "" &&
		body.Hcode == "" {

		tempData, err := service.repo.GetCradTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.CardData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	cardData, err := service.repo.GetCardData(dto.CardRequest{
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

	cidCountData, err := service.repo.GetCidCountData(dto.CardRequest{
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

	finalResult := []dto.CardData{
		calculateCardData(*cardData, *cidCountData),
	}

	return finalResult, nil
}

func calculateCardData(cardRawData dao.CardRawData, cidCountData dao.CidCountData) dto.CardData {
	var avgService float64
	var avgExpense float64
	expense := float64(cardRawData.DmExpense + cardRawData.HtExpense + cardRawData.CaExpense + cardRawData.CopdExpense + cardRawData.PsyExpense + cardRawData.HdCvdExpense)
	if cidCountData.CidCount > 0 {
		avgService = float64(cardRawData.ServiceCount) / float64(cidCountData.CidCount)
		avgExpense = expense / float64(cidCountData.CidCount)
	}

	return dto.CardData{
		ServiceCount: float64(cardRawData.ServiceCount),
		PatientCount: float64(cidCountData.CidCount),
		Expense:      utils.RoundToTwoDecimalPlaces(expense),
		HcodeCount:   cardRawData.HcodeCount,
		AvgService:   utils.RoundToTwoDecimalPlaces(avgService),
		AvgExpense:   utils.RoundToTwoDecimalPlaces(avgExpense),
	}
}
