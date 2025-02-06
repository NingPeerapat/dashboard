package dao

type ChartExpenseTempData struct {
	ChartExpenseTemp []TempExpenseData `bson:"chart_expense" json:"chart_expense"`
}

type TempExpenseData struct {
	DiseaseName  string  `bson:"diseaseName" json:"diseaseName"`
	QtyOfExpense float64 `bson:"qtyOfExpense" json:"qtyOfExpense"`
	Avg          float64 `bson:"avg" json:"avg"`
}
