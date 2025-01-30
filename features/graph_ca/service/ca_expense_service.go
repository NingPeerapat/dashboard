package service

import (
	"fmt"
	"ning/go-dashboard/features/graph_ca/entities"
	"ning/go-dashboard/features/graph_ca/repository"
	"ning/go-dashboard/pkg/utils"
	"strconv"
)

type GraphCaExpenseService struct {
	repo *repository.GraphCaExpenseRepository
}

func NewGraphCaExpenseService(repo *repository.GraphCaExpenseRepository) *GraphCaExpenseService {
	return &GraphCaExpenseService{repo: repo}
}

func (service *GraphCaExpenseService) GetExpenseData(body entities.CaRequest) ([]entities.CaData, error) {

	fullMonths := utils.GenerateFullMonths(body.Year)

	expenseData, err := service.repo.GetCaExpense(body)
	if err != nil {
		return nil, err
	}

	finalResult := FillPatientResults(fullMonths, expenseData)

	return finalResult, nil

}

func FillPatientResults(fullMonths []utils.FullMonthData, expenseData []entities.CaExpenseData) []entities.CaData {

	diseaseOrder := []string{
		"โรคมะเร็งปอด",
		"โรคมะเร็งเต้านม",
		"โรคมะเร็งปากมดลูก",
		"โรคมะเร็งตับ",
		"โรคมะเร็งลำไส้ใหญ่",
		"โรคมะเร็งชนิดอื่น ๆ",
	}

	diseaseMap := map[string][]float64{
		"โรคมะเร็งปอด":        {},
		"โรคมะเร็งเต้านม":     {},
		"โรคมะเร็งปากมดลูก":   {},
		"โรคมะเร็งตับ":        {},
		"โรคมะเร็งลำไส้ใหญ่":  {},
		"โรคมะเร็งชนิดอื่น ๆ": {},
	}

	for _, fm := range fullMonths {

		var found entities.CaExpenseData
		foundMatch := false
		for _, e := range expenseData {
			if e.Year == fm.Year && e.Month == fm.Month {
				found = e
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			found = entities.CaExpenseData{
				CaExpense:           0,
				LungCaExpense:       0,
				BreastCaExpense:     0,
				CervicalCaExpense:   0,
				LiverCaExpense:      0,
				ColorectalCaExpense: 0,
			}
		}

		otherCaExpense := found.CaExpense - (found.LungCaExpense + found.BreastCaExpense + found.CervicalCaExpense + found.LiverCaExpense + found.ColorectalCaExpense)

		diseaseMap["โรคมะเร็งปอด"] = append(diseaseMap["โรคมะเร็งปอด"], roundToTwoDecimal(found.LungCaExpense))
		diseaseMap["โรคมะเร็งเต้านม"] = append(diseaseMap["โรคมะเร็งเต้านม"], roundToTwoDecimal(found.BreastCaExpense))
		diseaseMap["โรคมะเร็งปากมดลูก"] = append(diseaseMap["โรคมะเร็งปากมดลูก"], roundToTwoDecimal(found.CervicalCaExpense))
		diseaseMap["โรคมะเร็งตับ"] = append(diseaseMap["โรคมะเร็งตับ"], roundToTwoDecimal(found.LiverCaExpense))
		diseaseMap["โรคมะเร็งลำไส้ใหญ่"] = append(diseaseMap["โรคมะเร็งลำไส้ใหญ่"], roundToTwoDecimal(found.ColorectalCaExpense))
		diseaseMap["โรคมะเร็งชนิดอื่น ๆ"] = append(diseaseMap["โรคมะเร็งชนิดอื่น ๆ"], roundToTwoDecimal(otherCaExpense))
	}

	var result []entities.CaData
	for _, diseaseName := range diseaseOrder {
		result = append(result, entities.CaData{
			DiseaseName: diseaseName,
			Data:        diseaseMap[diseaseName],
		})
	}

	return result
}

func roundToTwoDecimal(value float64) float64 {
	roundedValue, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return roundedValue
}
