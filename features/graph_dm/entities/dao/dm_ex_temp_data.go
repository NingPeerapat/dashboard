package dao

type GraphDmExTempData struct {
	GraphDmExpenseTemp []DmExTempData `bson:"dm_expense" json:"dm_expense"`
}

type DmExTempData struct {
	DiseaseName string    `bson:"diseaseName" json:"diseaseName"`
	Data        []float64 `bson:"data" json:"data"`
}
