package dto

type CardResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Result  []CardData `json:"result"`
}

type CardData struct {
	ServiceCount float64 `json:"service_count"`
	PatientCount float64 `json:"patient_count"`
	Expense      float64 `json:"expense"`
	HcodeCount   float64 `json:"hcode_count"`
	AvgService   float64 `json:"avg_service"`
	AvgExpense   float64 `json:"avg_expense"`
}
