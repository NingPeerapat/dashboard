package service

import (
	"log"
	"ning/go-dashboard/features/graph_ca/entities/dto"
	"ning/go-dashboard/features/graph_ca/repository"
	"ning/go-dashboard/pkg/utils"
	"sync"
)

type GraphCaPtService struct {
	repo *repository.GraphCaPtRepo
}

func NewGraphCaPtService(repo *repository.GraphCaPtRepo) *GraphCaPtService {
	return &GraphCaPtService{repo: repo}
}

func (service *GraphCaPtService) GetGraphCaPtData(body dto.CaRequest) ([]dto.CaData, error) {

	if body.Year == 2024 && body.Area == "" &&
		body.Province == "" &&
		body.District == "" &&
		body.Hcode == "" {

		tempData, err := service.repo.GetGraphCaPtTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.CaData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	var wg sync.WaitGroup
	cidData := make(map[string][]utils.PatientData)
	errChan := make(chan error, 6)
	var mu sync.Mutex

	callRepo := func(key string, fn func(body dto.CaRequest) ([]utils.PatientData, error)) {
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

	wg.Add(6)

	go callRepo("ca", service.repo.GetCaPatient)
	go callRepo("lung_ca", service.repo.GetLungCaPatient)
	go callRepo("breast_ca", service.repo.GetBreastCaPatient)
	go callRepo("cervical_ca", service.repo.GetCervicalCaPatient)
	go callRepo("liver_ca", service.repo.GetLiverCaPatient)
	go callRepo("colorectal_ca", service.repo.GetColorectalCaPatient)

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		for err := range errChan {
			log.Println("Error:", err)
		}
		return nil, <-errChan
	}

	fullMonths := utils.GenerateFullMonths(body.Year)

	caData := utils.FillPatientResults(fullMonths, cidData["ca"])
	lungCaData := utils.FillPatientResults(fullMonths, cidData["lung_ca"])
	breastCaData := utils.FillPatientResults(fullMonths, cidData["breast_ca"])
	cervicalCaData := utils.FillPatientResults(fullMonths, cidData["cervical_ca"])
	liverCaData := utils.FillPatientResults(fullMonths, cidData["liver_ca"])
	colorectalCaData := utils.FillPatientResults(fullMonths, cidData["colorectal_ca"])

	finalCaData := make([]float64, len(caData))

	for i, item := range caData {
		finalCaData[i] = item - (lungCaData[i] + breastCaData[i] + cervicalCaData[i] + liverCaData[i] + colorectalCaData[i])
	}

	finalResult := []dto.CaData{
		{
			DiseaseName: "โรคมะเร็งปอด",
			Data:        lungCaData,
		},
		{
			DiseaseName: "โรคมะเร็งเต้านม",
			Data:        breastCaData,
		},
		{
			DiseaseName: "โรคมะเร็งปากมดลูก",
			Data:        cervicalCaData,
		},
		{
			DiseaseName: "โรคมะเร็งตับ",
			Data:        liverCaData,
		},
		{
			DiseaseName: "โรคมะเร็งลำไส้ใหญ่",
			Data:        colorectalCaData,
		},
		{
			DiseaseName: "โรคมะเร็งชนิดอื่น ๆ",
			Data:        finalCaData,
		},
	}

	return finalResult, nil

}
