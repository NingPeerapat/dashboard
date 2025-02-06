package dao

type CardTempData struct {
	CardData []TempData `bson:"card_summary" json:"card_summary"`
}

type TempData struct {
	ServiceCount float64 `bson:"service_count" json:"service_count"`
	PatientCount float64 `bson:"patient_count" json:"patient_count"`
	Expense      float64 `bson:"expense" json:"expense"`
	HcodeCount   float64 `bson:"hcode_count" json:"hcode_count"`
	AvgService   float64 `bson:"avg_service" json:"avg_service"`
	AvgExpense   float64 `bson:"avg_expense" json:"avg_expense"`
}
