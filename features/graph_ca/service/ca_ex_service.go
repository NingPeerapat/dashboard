package service

import (
	"ning/go-dashboard/features/graph_ca/entities/dto"
	"ning/go-dashboard/features/graph_ca/repository"
	"ning/go-dashboard/pkg/utils"
)

type GraphCaExService struct {
	repo *repository.GraphCaExRepo
}

func NewGraphCaExService(repo *repository.GraphCaExRepo) *GraphCaExService {
	return &GraphCaExService{repo: repo}
}

func (service *GraphCaExService) GetGraphCaExData(body dto.CaRequest) ([]dto.CaData, error) {

	if body.Year == 2024 && body.Area == "" &&
		body.Province == "" &&
		body.District == "" &&
		body.Hcode == "" {

		tempData, err := service.repo.GetGraphCaExTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.CaData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	fullMonths := utils.GenerateFullMonths(body.Year)

	expenseData, err := service.repo.GetGraphCaExData(body)
	if err != nil {
		return nil, err
	}

	caRawExData := []utils.ExpenseData{}
	lungCaRawExData := []utils.ExpenseData{}
	breastCaRawExData := []utils.ExpenseData{}
	cervicalCaRawExData := []utils.ExpenseData{}
	liverCaRawExData := []utils.ExpenseData{}
	colorectalCaRawExData := []utils.ExpenseData{}

	for i := 0; i < len(expenseData); i++ {
		caRawExData = append(caRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].CaExpense,
		})
		lungCaRawExData = append(lungCaRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].LungCaExpense,
		})
		breastCaRawExData = append(breastCaRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].BreastCaExpense,
		})
		cervicalCaRawExData = append(cervicalCaRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].CervicalCaExpense,
		})
		liverCaRawExData = append(liverCaRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].LiverCaExpense,
		})
		colorectalCaRawExData = append(colorectalCaRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].ColorectalCaExpense,
		})
	}

	caExData := utils.FillExpenseResults(fullMonths, caRawExData)
	lungCaExData := utils.FillExpenseResults(fullMonths, lungCaRawExData)
	breastCaExData := utils.FillExpenseResults(fullMonths, breastCaRawExData)
	cervicalCaExData := utils.FillExpenseResults(fullMonths, cervicalCaRawExData)
	liverCaExData := utils.FillExpenseResults(fullMonths, liverCaRawExData)
	colorectalCaExData := utils.FillExpenseResults(fullMonths, colorectalCaRawExData)

	otherCaExData := make([]float64, len(caExData))

	for i := 0; i < 11; i++ {
		otherCaExData[i] = utils.RoundToTwoDecimalPlaces(caExData[i] - (lungCaExData[i] + breastCaExData[i] + cervicalCaExData[i] + liverCaExData[i] + colorectalCaExData[i]))
	}

	finalResult := []dto.CaData{
		{
			DiseaseName: "โรคมะเร็งปอด",
			Data:        lungCaExData,
		},
		{
			DiseaseName: "โรคมะเร็งเต้านม",
			Data:        breastCaExData,
		},
		{
			DiseaseName: "โรคมะเร็งปากมดลูก",
			Data:        cervicalCaExData,
		},
		{
			DiseaseName: "โรคมะเร็งตับ",
			Data:        liverCaExData,
		},
		{
			DiseaseName: "โรคมะเร็งลำไส้ใหญ่",
			Data:        colorectalCaExData,
		},
		{
			DiseaseName: "โรคมะเร็งชนิดอื่น ๆ",
			Data:        otherCaExData,
		},
	}

	return finalResult, nil
}
