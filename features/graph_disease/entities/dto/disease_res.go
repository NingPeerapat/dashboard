package dto

type DiseaseResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Result  []DiseaseData `json:"result"`
}

type DiseaseData struct {
	DiseaseName string    `json:"diseaseName"`
	Data        []float64 `json:"data"`
}
