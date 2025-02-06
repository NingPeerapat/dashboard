package dao

type GraphDiseasePtTempData struct {
	GraphDiseasePtpenseTemp []DiseasePtTempData `bson:"disease_patient" json:"disease_patient"`
}

type DiseasePtTempData struct {
	DiseaseName string    `bson:"diseaseName" json:"diseaseName"`
	Data        []float64 `bson:"data" json:"data"`
}
