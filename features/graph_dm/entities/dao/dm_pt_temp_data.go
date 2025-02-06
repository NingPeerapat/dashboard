package dao

type GraphDmPtTempData struct {
	GraphDmPtpenseTemp []DmPtTempData `bson:"dm_patient" json:"dm_patient"`
}

type DmPtTempData struct {
	DiseaseName string    `bson:"diseaseName" json:"diseaseName"`
	Data        []float64 `bson:"data" json:"data"`
}
