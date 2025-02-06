package dto

type ChartExpenseResponse struct {
	Status  bool               `json:"status"`
	Message string             `json:"message"`
	Result  []ChartExpenseData `json:"result"`
}

type ChartExpenseData struct {
	DiseaseName  string  `json:"diseaseName"`
	QtyOfExpense float64 `json:"qtyOfExpense"`
	Avg          float64 `json:"avg"`
}
