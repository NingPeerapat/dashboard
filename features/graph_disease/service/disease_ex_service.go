package service

import (
	"ning/go-dashboard/features/graph_disease/entities/dto"
	"ning/go-dashboard/features/graph_disease/repository"
	"ning/go-dashboard/pkg/utils"
)

type GraphDiseaseExService struct {
	repo *repository.GraphDiseaseExRepo
}

func NewGraphDiseaseExService(repo *repository.GraphDiseaseExRepo) *GraphDiseaseExService {
	return &GraphDiseaseExService{repo: repo}
}

func (service *GraphDiseaseExService) GetGraphDiseaseExData(body dto.DiseaseRequest) ([]dto.DiseaseData, error) {

	fullMonths := utils.GenerateFullMonths(body.Year)

	expenseData, err := service.repo.GetGraphDiseaseExData(body)
	if err != nil {
		return nil, err
	}

	dmRawExData := []utils.ExpenseData{}
	htRawExData := []utils.ExpenseData{}
	copdRawExData := []utils.ExpenseData{}
	caRawExData := []utils.ExpenseData{}
	psyRawExData := []utils.ExpenseData{}
	hdCvdRawExData := []utils.ExpenseData{}

	for i := 0; i < len(expenseData); i++ {
		dmRawExData = append(dmRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].DmExpense,
		})
		htRawExData = append(htRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].HtExpense,
		})
		copdRawExData = append(copdRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].CopdExpense,
		})
		caRawExData = append(caRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].CaExpense,
		})
		psyRawExData = append(psyRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].PsyExpense,
		})
		hdCvdRawExData = append(hdCvdRawExData, utils.ExpenseData{
			Year:    expenseData[i].Year,
			Month:   expenseData[i].Month,
			Expense: expenseData[i].HdCvdExpense,
		})
	}

	dmExData := utils.FillExpenseResults(fullMonths, dmRawExData)
	htExData := utils.FillExpenseResults(fullMonths, htRawExData)
	copdExData := utils.FillExpenseResults(fullMonths, copdRawExData)
	caExData := utils.FillExpenseResults(fullMonths, caRawExData)
	psyExData := utils.FillExpenseResults(fullMonths, psyRawExData)
	hdCvdExData := utils.FillExpenseResults(fullMonths, hdCvdRawExData)

	allDiseaseExData := make([]float64, len(dmExData))

	for i := 0; i < 11; i++ {
		allDiseaseExData[i] = utils.RoundToTwoDecimalPlaces(dmExData[i] + htExData[i] + copdExData[i] + caExData[i] + psyExData[i] + hdCvdExData[i])
	}

	finalResult := []dto.DiseaseData{
		{
			DiseaseName: "ทั้งหมด",
			Data:        allDiseaseExData,
		},
		{
			DiseaseName: "โรคเบาหวาน",
			Data:        dmExData,
		},
		{
			DiseaseName: "โรคความดันโลหิตสูง",
			Data:        htExData,
		},
		{
			DiseaseName: "โรคปอดอุดกั้นเรื้อรัง",
			Data:        copdExData,
		},
		{
			DiseaseName: "โรคมะเร็ง",
			Data:        caExData,
		},
		{
			DiseaseName: "สุขภาพจิต",
			Data:        psyExData,
		},
		{
			DiseaseName: "โรคหัวใจและหลอดเลือด",
			Data:        hdCvdExData,
		},
	}

	return finalResult, nil
}
