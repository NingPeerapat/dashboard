package dto

type ChartPatientResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Result  []ChartPatientData `json:"result"`
}

type ChartPatientData struct {
	DiseaseName  string  `json:"diseaseName"`
	QtyOfPatient int     `json:"qtyOfPatient"`
	Avg          float64 `json:"avg"`
}
