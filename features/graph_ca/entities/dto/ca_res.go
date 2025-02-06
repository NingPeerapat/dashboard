package dto

type CaResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Result  []CaData `json:"result"`
}

type CaData struct {
	DiseaseName string    `json:"diseaseName"`
	Data        []float64 `json:"data"`
}
