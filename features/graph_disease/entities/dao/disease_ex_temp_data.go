package dao

type GraphDiseaseExTempData struct {
	GraphDiseaseExpenseTemp []DiseaseExTempData `bson:"disease_expense" json:"disease_expense"`
}

type DiseaseExTempData struct {
	DiseaseName string    `bson:"diseaseName" json:"diseaseName"`
	Data        []float64 `bson:"data" json:"data"`
}
