package dto

type DmResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Result  []DmData `json:"result"`
}

type DmData struct {
	DiseaseName string    `json:"diseaseName"`
	Data        []float64 `json:"data"`
}
