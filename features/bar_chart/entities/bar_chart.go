package entities

import "time"

type ChartCilent struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Area      string `json:"area"`
	Province  string `json:"province"`
	District  string `json:"district"`
	Hcode     string `json:"hcode"`
}

type ChartRequest struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Area      string    `json:"area"`
	Province  string    `json:"province"`
	District  string    `json:"district"`
	Hcode     string    `json:"hcode"`
}

type ChartExpenseResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Result  []DiseaseExpenseData `json:"result"`
}

type DiseaseExpenseData struct {
	DiseaseName  string  `json:"diseaseName"`
	QtyOfExpense float64 `json:"qtyOfExpense"`
	Avg          float64 `json:"avg"`
}

type ExpenseData struct {
	DmExpense    float64 `bson:"dm_expense" json:"dm_expense"`
	HtExpense    float64 `bson:"ht_expense" json:"ht_expense"`
	CopdExpense  float64 `bson:"copd_expense" json:"copd_expense"`
	CaExpense    float64 `bson:"ca_expense" json:"ca_expense"`
	PsyExpense   float64 `bson:"psy_expense" json:"psy_expense"`
	HdCvdExpense float64 `bson:"hd_cvd_expense" json:"hd_cvd_expense"`
}

type ChartPatientResponse struct {
	Status  bool                 `json:"status"`
	Message string               `json:"message"`
	Result  []DiseasePatientData `json:"result"`
}

type DiseasePatientData struct {
	DiseaseName  string  `json:"diseaseName"`
	QtyOfPatient int     `json:"qtyOfPatient"`
	Avg          float64 `json:"avg"`
}

type UidData struct {
	DmUidCount    float64 `bson:"dm_uid_count" json:"dm_uid_count"`
	HtUidCount    float64 `bson:"ht_uid_count" json:"ht_uid_count"`
	CopdUidCount  float64 `bson:"copd_uid_count" json:"copd_uid_count"`
	CaUidCount    float64 `bson:"ca_uid_count" json:"ca_uid_count"`
	PsyUidCount   float64 `bson:"psy_uid_count" json:"psy_uid_count"`
	HdCvdUidCount float64 `bson:"hd_cvd_uid_count" json:"hd_cvd_uid_count"`
}
