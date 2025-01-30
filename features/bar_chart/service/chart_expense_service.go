package service

import (
	"fmt"
	"log"
	"ning/go-dashboard/features/bar_chart/entities"
	"ning/go-dashboard/features/bar_chart/repository"
	"ning/go-dashboard/pkg/shared"
	"ning/go-dashboard/pkg/utils"
	"sync"
)

type ChartExpenseService struct {
	repo  *repository.ChartExpenseRepository
	count *shared.CountCidRepository
}

func NewChartExpenseService(repo *repository.ChartExpenseRepository, count *shared.CountCidRepository) *ChartExpenseService {
	return &ChartExpenseService{repo: repo, count: count}
}

func (service *ChartExpenseService) GetExpenseData(body entities.ChartRequest) ([]entities.DiseaseExpenseData, error) {

	var wg sync.WaitGroup
	countData := make(map[string]int)
	expenseChan := make(chan interface{}, 1)
	errChan := make(chan error, 6)
	var mu sync.Mutex

	callRepo := func(key string, fn func(entities.ChartRequest) (int, error)) {
		defer wg.Done()
		count, err := fn(body)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		countData[key] = count
		mu.Unlock()
	}

	wg.Add(7)

	go callRepo("dm", service.count.CountDmCid)
	go callRepo("ht", service.count.CountHtCid)
	go callRepo("copd", service.count.CountCopdCid)
	go callRepo("ca", service.count.CountCaCid)
	go callRepo("psy", service.count.CountPsyCid)
	go callRepo("hd_cvd", service.count.CountHdCvdCid)

	go func() {
		defer wg.Done()
		data, err := service.repo.GetAllData(entities.ChartRequest{
			StartDate: body.StartDate,
			EndDate:   body.EndDate,
			Area:      body.Area,
			Province:  body.Province,
			District:  body.District,
			Hcode:     body.Hcode,
		})
		if err != nil {
			errChan <- err
			return
		}

		expenseChan <- data

	}()

	wg.Wait()
	close(errChan)
	close(expenseChan)

	var finalErr error
	for err := range errChan {
		log.Println("Error:", err)
		finalErr = err
	}
	if finalErr != nil {
		return nil, finalErr
	}

	expenseDataInterface := <-expenseChan

	expenseData, ok := expenseDataInterface.(*entities.ExpenseData)
	if !ok {
		return nil, fmt.Errorf("failed to cast expenseData to entities.ExpenseData, received: %+v", expenseDataInterface)
	}

	finalResult := []entities.DiseaseExpenseData{
		calculateDiseaseExpenseData("โรคเบาหวาน", expenseData.DmExpense, countData["dm"]),
		calculateDiseaseExpenseData("โรคความดันโลหิตสูง", expenseData.HtExpense, countData["ht"]),
		calculateDiseaseExpenseData("โรคปอดอุดกั้นเรื้อรัง", expenseData.CopdExpense, countData["copd"]),
		calculateDiseaseExpenseData("โรคมะเร็ง", expenseData.CaExpense, countData["ca"]),
		calculateDiseaseExpenseData("สุขภาพจิต", expenseData.PsyExpense, countData["psy"]),
		calculateDiseaseExpenseData("โรคหัวใจและหลอดเลือด", expenseData.HdCvdExpense, countData["hd_cvd"]),
	}

	return finalResult, nil
}

func calculateDiseaseExpenseData(diseaseName string, expenseKey float64, cidCountKey int) entities.DiseaseExpenseData {
	var avg float64
	if cidCountKey > 0 {
		avg = expenseKey / float64(cidCountKey)
	}

	return entities.DiseaseExpenseData{
		DiseaseName:  diseaseName,
		QtyOfExpense: utils.RoundToTwoDecimalPlaces(expenseKey),
		Avg:          utils.RoundToTwoDecimalPlaces(avg),
	}
}
