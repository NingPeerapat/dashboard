package dto

type ChartPatientResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Result  []ChartPatientData `json:"result"`
}

type ChartPatientData struct {
	DiseaseName  string  `bson:"diseaseName" json:"diseaseName"`
	QtyOfPatient int     `bson:"qtyOfPatient" json:"qtyOfPatient"`
	Avg          float64 `bson:"avg" json:"avg"`
}
