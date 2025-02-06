package dao

type GraphCaExTempData struct {
	GraphCaExpenseTemp []CaExTempData `bson:"ca_expense" json:"ca_expense"`
}

type CaExTempData struct {
	DiseaseName string    `bson:"diseaseName" json:"diseaseName"`
	Data        []float64 `bson:"data" json:"data"`
}
