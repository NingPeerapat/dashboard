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

type ChartPtService struct {
	repo  *repository.ChartPtRepo
	count *repository.CountCidRepo
}

func NewChartPtService(repo *repository.ChartPtRepo, count *repository.CountCidRepo) *ChartPtService {
	return &ChartPtService{repo: repo, count: count}
}

func (service *ChartPtService) GetChartPtData(body dto.ChartRequest) ([]dto.ChartPatientData, error) {

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

		tempData, err := service.repo.GetChartPtTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.ChartPatientData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	var wg sync.WaitGroup
	countData := make(map[string]int)
	uidChan := make(chan interface{}, 1)
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
		data, err := service.repo.GetChartUidData(dto.ChartRequest{
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

		uidChan <- data

	}()

	wg.Wait()
	close(errChan)
	close(uidChan)

	if len(errChan) > 0 {
		for err := range errChan {
			log.Println("Error:", err)
		}
		return nil, <-errChan
	}
	uidData, ok := (<-uidChan).(*dao.UidData)
	if !ok {
		return nil, fmt.Errorf("failed to cast uidData to ChartPatientData")
	}

	finalResult := []dto.ChartPatientData{
		calculateChartPtData("โรคเบาหวาน", int(uidData.DmUidCount), countData["dm"]),
		calculateChartPtData("โรคความดันโลหิตสูง", int(uidData.HtUidCount), countData["ht"]),
		calculateChartPtData("โรคปอดอุดกั้นเรื้อรัง", int(uidData.CopdUidCount), countData["copd"]),
		calculateChartPtData("โรคมะเร็ง", int(uidData.CaUidCount), countData["ca"]),
		calculateChartPtData("สุขภาพจิต", int(uidData.PsyUidCount), countData["psy"]),
		calculateChartPtData("โรคหัวใจและหลอดเลือด", int(uidData.HdCvdUidCount), countData["hd_cvd"]),
	}

	return finalResult, nil
}

func calculateChartPtData(diseaseName string, uidKey int, cidCountKey int) dto.ChartPatientData {
	var avg float64
	if cidCountKey > 0 {
		avg = float64(uidKey) / float64(cidCountKey)
	}

	return dto.ChartPatientData{
		DiseaseName:  diseaseName,
		QtyOfPatient: int(uidKey),
		Avg:          utils.RoundToTwoDecimalPlaces(avg),
	}
}
