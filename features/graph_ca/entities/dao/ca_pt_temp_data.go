package dao

type GraphCaPtTempData struct {
	GraphCaPtpenseTemp []CaPtTempData `bson:"ca_patient" json:"ca_patient"`
}

type CaPtTempData struct {
	DiseaseName string    `bson:"diseaseName" json:"diseaseName"`
	Data        []float64 `bson:"data" json:"data"`
}
