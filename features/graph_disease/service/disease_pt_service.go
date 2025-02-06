package service

import (
	"log"
	"ning/go-dashboard/features/graph_disease/entities/dto"
	"ning/go-dashboard/features/graph_disease/repository"
	"ning/go-dashboard/pkg/utils"
	"sync"
)

type GraphDiseasePtService struct {
	repo *repository.GraphDiseasePtRepo
}

func NewGraphDiseasePtService(repo *repository.GraphDiseasePtRepo) *GraphDiseasePtService {
	return &GraphDiseasePtService{repo: repo}
}

func (service *GraphDiseasePtService) GetGraphDiseasePtData(body dto.DiseaseRequest) ([]dto.DiseaseData, error) {

	if body.Year == 2024 && body.Area == "" &&
		body.Province == "" &&
		body.District == "" &&
		body.Hcode == "" {

		tempData, err := service.repo.GetGraphDiseasePtTempData()
		if err != nil {
			return nil, err
		}
		convertedData := make([]dto.DiseaseData, len(tempData))
		for i, v := range tempData {
			convertedData[i] = *v
		}
		return convertedData, nil
	}

	var wg sync.WaitGroup
	cidData := make(map[string][]utils.PatientData)
	errChan := make(chan error, 6)
	var mu sync.Mutex

	callRepo := func(key string, fn func(body dto.DiseaseRequest) ([]utils.PatientData, error)) {
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

	go callRepo("dm", service.repo.GetDmPatient)
	go callRepo("ht", service.repo.GetHtPatient)
	go callRepo("copd", service.repo.GetCopdPatient)
	go callRepo("ca", service.repo.GetCaPatient)
	go callRepo("psy", service.repo.GetPsyPatient)
	go callRepo("hd_cvd", service.repo.GetHdCvdPatient)

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		for err := range errChan {
			log.Println("Error:", err)
		}
		return nil, <-errChan
	}

	fullMonths := utils.GenerateFullMonths(body.Year)

	dmData := utils.FillPatientResults(fullMonths, cidData["dm"])
	htData := utils.FillPatientResults(fullMonths, cidData["ht"])
	copdData := utils.FillPatientResults(fullMonths, cidData["copd"])
	caData := utils.FillPatientResults(fullMonths, cidData["ca"])
	psyData := utils.FillPatientResults(fullMonths, cidData["psy"])
	hdCvdData := utils.FillPatientResults(fullMonths, cidData["hd_cvd"])

	allDiseaseData := make([]float64, len(dmData))

	for i := 0; i < 11; i++ {
		allDiseaseData[i] = dmData[i] + htData[i] + copdData[i] + caData[i] + psyData[i] + hdCvdData[i]
	}

	finalResult := []dto.DiseaseData{
		{
			DiseaseName: "ทั้งหมด",
			Data:        allDiseaseData,
		},
		{
			DiseaseName: "โรคเบาหวาน",
			Data:        dmData,
		},
		{
			DiseaseName: "โรคความดันโลหิตสูง",
			Data:        htData,
		},
		{
			DiseaseName: "โรคปอดอุดกั้นเรื้อรัง",
			Data:        copdData,
		},
		{
			DiseaseName: "โรคมะเร็ง",
			Data:        caData,
		},
		{
			DiseaseName: "สุขภาพจิต",
			Data:        psyData,
		},
		{
			DiseaseName: "โรคหัวใจและหลอดเลือด",
			Data:        hdCvdData,
		},
	}

	return finalResult, nil

}
