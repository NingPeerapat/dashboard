package service

import (
	"log"
	"ning/go-dashboard/features/graph_dm/entities/dto"
	"ning/go-dashboard/features/graph_dm/repository"
	"ning/go-dashboard/pkg/utils"
	"sync"
)

type GraphDmPtService struct {
	repo *repository.GraphDmPtRepo
}

func NewGraphDmPtService(repo *repository.GraphDmPtRepo) *GraphDmPtService {
	return &GraphDmPtService{repo: repo}
}

func (service *GraphDmPtService) GetGraphDmPtData(body dto.DmRequest) ([]dto.DmData, error) {

	if body.Year == 2024 && body.Area == "" &&
		body.Province == "" &&
		body.District == "" &&
		body.Hcode == "" {

		tempData, err := service.repo.GetGraphDmPtTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.DmData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	var wg sync.WaitGroup
	cidData := make(map[string][]utils.PatientData)
	errChan := make(chan error, 6)
	var mu sync.Mutex

	callRepo := func(key string, fn func(body dto.DmRequest) ([]utils.PatientData, error)) {
		defer wg.Done()
		count, err := fn(body)
		if err != nil {
			errChan <- err
			return
		}
		mu.Lock()
		cidData[key] = count
		mu.Unlock()
	}

	wg.Add(5)

	go callRepo("dm", service.repo.GetDmPatient)
	go callRepo("hg", service.repo.GetHgPatient)
	go callRepo("dm_ckd", service.repo.GetDmCkdPatient)
	go callRepo("dm_acs", service.repo.GetDmAcsPatient)
	go callRepo("dm_cva", service.repo.GetDmCvaPatient)

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		for err := range errChan {
			log.Println("Error:", err)
		}
		return nil, <-errChan
	}

	fullMonths := utils.GenerateFullMonths(body.Year)

	dmPtData := utils.FillPatientResults(fullMonths, cidData["dm"])
	hgPtData := utils.FillPatientResults(fullMonths, cidData["hg"])
	dmCkdPtData := utils.FillPatientResults(fullMonths, cidData["dm_ckd"])
	dmAcsPtData := utils.FillPatientResults(fullMonths, cidData["dm_acs"])
	dmCvaPtData := utils.FillPatientResults(fullMonths, cidData["dm_cva"])

	otherDmPtData := make([]float64, len(dmPtData))

	for i := 0; i < 11; i++ {
		otherDmPtData[i] = dmPtData[i] - (hgPtData[i] + dmCkdPtData[i] + dmAcsPtData[i] + dmCvaPtData[i])
	}

	finalResult := []dto.DmData{
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ที่มีภาวะน้ำตาลในเลือดสูง",
			Data:        hgPtData,
		},
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ร่วมกับโรคไตเรื้อรัง",
			Data:        dmCkdPtData,
		},
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ร่วมกับโรคหัวใจขาดเลือด",
			Data:        dmAcsPtData,
		},
		{
			DiseaseName: "เบาหวานชนิดต่าง ๆ ร่วมกับโรคหลอดเลือดสมอง",
			Data:        dmCvaPtData,
		},
		{
			DiseaseName: "เบาหวานชนิดอื่น ๆ",
			Data:        otherDmPtData,
		},
	}

	return finalResult, nil

}
