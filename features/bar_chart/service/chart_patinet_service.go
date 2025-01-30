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

type ChartPatientService struct {
	repo  *repository.ChartPatientRepository
	count *shared.CountCidRepository
}

func NewChartPatientService(repo *repository.ChartPatientRepository, count *shared.CountCidRepository) *ChartPatientService {
	return &ChartPatientService{repo: repo, count: count}
}

func (service *ChartPatientService) GetPatientData(body entities.ChartRequest) ([]entities.DiseasePatientData, error) {

	var wg sync.WaitGroup
	countData := make(map[string]int)
	uidChan := make(chan interface{}, 1)
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
	uidData, ok := (<-uidChan).(*entities.UidData)
	if !ok {
		return nil, fmt.Errorf("failed to cast uidData to ChartPatientData")
	}

	finalResult := []entities.DiseasePatientData{
		calculateDiseasePatientData("โรคเบาหวาน", int(uidData.DmUidCount), countData["dm"]),
		calculateDiseasePatientData("โรคความดันโลหิตสูง", int(uidData.HtUidCount), countData["ht"]),
		calculateDiseasePatientData("โรคปอดอุดกั้นเรื้อรัง", int(uidData.CopdUidCount), countData["copd"]),
		calculateDiseasePatientData("โรคมะเร็ง", int(uidData.CaUidCount), countData["ca"]),
		calculateDiseasePatientData("สุขภาพจิต", int(uidData.PsyUidCount), countData["psy"]),
		calculateDiseasePatientData("โรคหัวใจและหลอดเลือด", int(uidData.HdCvdUidCount), countData["hd_cvd"]),
	}

	return finalResult, nil
}

func calculateDiseasePatientData(diseaseName string, uidKey int, cidCountKey int) entities.DiseasePatientData {
	var avg float64
	if cidCountKey > 0 {
		avg = float64(uidKey) / float64(cidCountKey)
	}

	return entities.DiseasePatientData{
		DiseaseName:  diseaseName,
		QtyOfPatient: int(uidKey),
		Avg:          utils.RoundToTwoDecimalPlaces(avg),
	}
}
