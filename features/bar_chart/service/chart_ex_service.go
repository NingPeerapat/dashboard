package service

import (
	"fmt"
	"log"
	"ning/go-dashboard/features/bar_chart/entities/dao"
	"ning/go-dashboard/features/bar_chart/entities/dto"
	"ning/go-dashboard/features/bar_chart/repository"
	"ning/go-dashboard/pkg/utils"
	"sync"
	"time"
)

type ChartExService struct {
	repo  *repository.ChartExRepo
	count *repository.CountCidRepo
}

func NewChartExService(repo *repository.ChartExRepo, count *repository.CountCidRepo) *ChartExService {
	return &ChartExService{repo: repo, count: count}
}

func (service *ChartExService) GetChartExData(body dto.ChartRequest) ([]dto.ChartExpenseData, error) {

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

		tempData, err := service.repo.GetChartExTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.ChartExpenseData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	var wg sync.WaitGroup
	countData := make(map[string]int)
	expenseChan := make(chan interface{}, 1)
	errChan := make(chan error, 6)
	var mu sync.Mutex

	callRepo := func(key string, fn func(dto.ChartRequest) (int, error)) {
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
		data, err := service.repo.GetChartExData(dto.ChartRequest{
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

	expenseData, ok := expenseDataInterface.(*dao.ExpenseData)
	if !ok {
		return nil, fmt.Errorf("failed to cast expenseData to entities ofExpenseData, received: %+v", expenseDataInterface)
	}

	finalResult := []dto.ChartExpenseData{
		calculateChartExData("โรคเบาหวาน", expenseData.DmExpense, countData["dm"]),
		calculateChartExData("โรคความดันโลหิตสูง", expenseData.HtExpense, countData["ht"]),
		calculateChartExData("โรคปอดอุดกั้นเรื้อรัง", expenseData.CopdExpense, countData["copd"]),
		calculateChartExData("โรคมะเร็ง", expenseData.CaExpense, countData["ca"]),
		calculateChartExData("สุขภาพจิต", expenseData.PsyExpense, countData["psy"]),
		calculateChartExData("โรคหัวใจและหลอดเลือด", expenseData.HdCvdExpense, countData["hd_cvd"]),
	}

	return finalResult, nil
}

func calculateChartExData(diseaseName string, expenseKey float64, cidCountKey int) dto.ChartExpenseData {
	var avg float64
	if cidCountKey > 0 {
		avg = expenseKey / float64(cidCountKey)
	}

	return dto.ChartExpenseData{
		DiseaseName:  diseaseName,
		QtyOfExpense: utils.RoundToTwoDecimalPlaces(expenseKey),
		Avg:          utils.RoundToTwoDecimalPlaces(avg),
	}
}
