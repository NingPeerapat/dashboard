package service

import (
	"fmt"
	"ning/go-dashboard/features/graph_disease/entities"
	"ning/go-dashboard/features/graph_disease/repository"
	"ning/go-dashboard/pkg/utils"
	"strconv"
)

type DiseaseExpenseService struct {
	repo *repository.DiseaseExpenseRepository
}

func NewDiseaseExpenseService(repo *repository.DiseaseExpenseRepository) *DiseaseExpenseService {
	return &DiseaseExpenseService{repo: repo}
}

func (service *DiseaseExpenseService) GetExpenseData(body entities.DiseaseRequest) ([]entities.DiseaseData, error) {

	fullMonths := utils.GenerateFullMonths(body.Year)

	expenseData, err := service.repo.GetDiseaseExpense(body)
	if err != nil {
		return nil, err
	}

	finalResult := FillPatientResults(fullMonths, expenseData)

	return finalResult, nil

}

func FillPatientResults(fullMonths []utils.FullMonthData, expenseData []entities.DiseaseExpenseData) []entities.DiseaseData {

	diseaseOrder := []string{
		"ทั้งหมด",
		"โรคเบาหวาน",
		"โรคความดันโลหิตสูง",
		"โรคปอดอุดกั้นเรื้อรัง",
		"โรคมะเร็ง",
		"สุขภาพจิต",
		"โรคหัวใจและหลอดเลือด",
	}

	diseaseMap := map[string][]float64{
		"ทั้งหมด":               {},
		"โรคเบาหวาน":            {},
		"โรคความดันโลหิตสูง":    {},
		"โรคปอดอุดกั้นเรื้อรัง": {},
		"โรคมะเร็ง":             {},
		"สุขภาพจิต":             {},
		"โรคหัวใจและหลอดเลือด":  {},
	}

	for _, fm := range fullMonths {

		var found entities.DiseaseExpenseData
		foundMatch := false
		for _, e := range expenseData {
			if e.Year == fm.Year && e.Month == fm.Month {
				found = e
				foundMatch = true
				break
			}
		}

		if !foundMatch {
			found = entities.DiseaseExpenseData{
				DMExpense:    0,
				HtExpense:    0,
				CopdExpense:  0,
				CaExpense:    0,
				PsyExpense:   0,
				HdCvdExpense: 0,
			}
		}

		allDiseaseData := found.DMExpense + found.HtExpense + found.CopdExpense + found.CaExpense + found.PsyExpense + found.HdCvdExpense

		diseaseMap["ทั้งหมด"] = append(diseaseMap["ทั้งหมด"], roundToTwoDecimal(allDiseaseData))
		diseaseMap["โรคเบาหวาน"] = append(diseaseMap["โรคเบาหวาน"], roundToTwoDecimal(found.DMExpense))
		diseaseMap["โรคความดันโลหิตสูง"] = append(diseaseMap["โรคความดันโลหิตสูง"], roundToTwoDecimal(found.HtExpense))
		diseaseMap["โรคปอดอุดกั้นเรื้อรัง"] = append(diseaseMap["โรคปอดอุดกั้นเรื้อรัง"], roundToTwoDecimal(found.CopdExpense))
		diseaseMap["โรคมะเร็ง"] = append(diseaseMap["โรคมะเร็ง"], roundToTwoDecimal(found.CaExpense))
		diseaseMap["สุขภาพจิต"] = append(diseaseMap["สุขภาพจิต"], roundToTwoDecimal(found.PsyExpense))
		diseaseMap["โรคหัวใจและหลอดเลือด"] = append(diseaseMap["โรคหัวใจและหลอดเลือด"], roundToTwoDecimal(found.HdCvdExpense))
	}

	var result []entities.DiseaseData
	for _, diseaseName := range diseaseOrder {
		result = append(result, entities.DiseaseData{
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
