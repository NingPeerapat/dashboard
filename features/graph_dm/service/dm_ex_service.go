package service

import (
	"ning/go-dashboard/features/graph_dm/entities/dto"
	"ning/go-dashboard/features/graph_dm/repository"
	"ning/go-dashboard/pkg/utils"
)

type GraphDmExService struct {
	repo *repository.GraphDmExRepo
}

func NewGraphDmExService(repo *repository.GraphDmExRepo) *GraphDmExService {
	return &GraphDmExService{repo: repo}
}

func (service *GraphDmExService) GetGraphDmExData(body dto.DmRequest) ([]dto.DmData, error) {

	if body.Year == 2024 && body.Area == "" &&
		body.Province == "" &&
		body.District == "" &&
		body.Hcode == "" {

		tempData, err := service.repo.GetGraphDiseaseExTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.DmData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	fullMonths := utils.GenerateFullMonths(body.Year)

	expenseData, err := service.repo.GetGraphDmExData(body)
	if err != nil {
		return nil, err
	}

	dmRawExData := []utils.ExpenseData{}
	hgRawExData := []utils.ExpenseData{}
	dmCkdRawExData := []utils.ExpenseData{}
	dmAcsRawExData := []utils.ExpenseData{}
	dmCvaRawExData := []utils.ExpenseData{}

	for i := 0; i < len(expenseData); i++ {
		dmRawExData = append(dmRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].DmExpense,
		})
		hgRawExData = append(hgRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].HgExpense,
		})
		dmCkdRawExData = append(dmCkdRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].DmCkdExpense,
		})
		dmAcsRawExData = append(dmAcsRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].DmAcsExpense,
		})
		dmCvaRawExData = append(dmCvaRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].DmCvaExpense,
		})
	}

	dmExData := utils.FillExpenseResults(fullMonths, dmRawExData)
	hgExData := utils.FillExpenseResults(fullMonths, hgRawExData)
	dmCkdExData := utils.FillExpenseResults(fullMonths, dmCkdRawExData)
	dmAcsExData := utils.FillExpenseResults(fullMonths, dmAcsRawExData)
	dmCvaExData := utils.FillExpenseResults(fullMonths, dmCvaRawExData)

	otherDmExData := make([]float64, len(dmExData))

	for i := 0; i < 11; i++ {
		otherDmExData[i] = utils.RoundToTwoDecimalPlaces(dmExData[i] - (hgExData[i] + dmCkdExData[i] + dmAcsExData[i] + dmCvaExData[i]))
	}

	finalResult := []dto.DmData{
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ที่มีภาวะน้ำตาลในเลือดสูง",
			Data:        hgExData,
		},
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ร่วมกับโรคไตเรื้อรัง",
			Data:        dmCkdExData,
		},
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ร่วมกับโรคหัวใจขาดเลือด",
			Data:        dmAcsExData,
		},
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ร่วมกับโรคหลอดเลือดสมอง",
			Data:        dmCvaExData,
		},
		{
			DiseaseName: "เบาหวานชนิดอื่น ๆ",
			Data:        otherDmExData,
		},
	}

	return finalResult, nil
}
