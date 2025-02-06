package dao

type DiseaseExpenseData struct {
	Year         int     `bson:"year" json:"year"`
	Month        int     `bson:"month" json:"month"`
	DmExpense    float64 `bson:"dm_expense" json:"dm_expense"`
	HtExpense    float64 `bson:"ht_expense" json:"ht_expense"`
	CopdExpense  float64 `bson:"copd_expense" json:"copd_expense"`
	CaExpense    float64 `bson:"ca_expense" json:"ca_expense"`
	PsyExpense   float64 `bson:"psy_expense" json:"psy_expense"`
	HdCvdExpense float64 `bson:"hd_cvd_expense" json:"hd_cvd_expense"`
}
