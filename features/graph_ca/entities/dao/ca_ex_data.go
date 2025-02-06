package dao

type CaExpenseData struct {
	Year                int     `bson:"year" json:"year"`
	Month               int     `bson:"month" json:"month"`
	CaExpense           float64 `bson:"ca_expense" json:"ca_expense"`
	LungCaExpense       float64 `bson:"lung_ca_expense" json:"lung_ca_expense"`
	BreastCaExpense     float64 `bson:"breast_ca_expense" json:"breast_ca_expense"`
	CervicalCaExpense   float64 `bson:"cervical_ca_expense" json:"cervical_ca_expense"`
	LiverCaExpense      float64 `bson:"liver_ca_expense" json:"liver_ca_expense"`
	ColorectalCaExpense float64 `bson:"colorectal_ca_expense" json:"colorectal_ca_expense"`
}
