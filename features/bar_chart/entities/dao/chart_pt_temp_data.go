package dao

type ChartPatientTempData struct {
	ChartPatientTemp []TempPatientData `bson:"chart_patient" json:"chart_patient"`
}

type TempPatientData struct {
	DiseaseName  string  `bson:"diseaseName" json:"diseaseName"`
	QtyOfPatient int     `bson:"qtyOfPatient" json:"qtyOfPatient"`
	Avg          float64 `bson:"avg" json:"avg"`
}
