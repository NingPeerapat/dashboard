package entities

import "time"

type CardCilent struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Area      string `json:"area"`
	Province  string `json:"province"`
	District  string `json:"district"`
	Hcode     string `json:"hcode"`
}

type CardRequest struct {
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Area      string    `json:"area"`
	Province  string    `json:"province"`
	District  string    `json:"district"`
	Hcode     string    `json:"hcode"`
}

type CardResponse struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Result  []CardData `json:"result"`
}

type CardData struct {
	ServiceCount float64 `bson:"service_count" json:"service_count"`
	PatientCount float64 `bson:"patient_count" json:"patient_count"`
	Expense      float64 `bson:"expense" json:"expense"`
	HcodeCount   float64 `bson:"hcode_count" json:"hcode_count"`
	AvgService   float64 `bson:"avg_service" json:"avg_service"`
	AvgExpense   float64 `bson:"avg_expense" json:"avg_expense"`
}

type CardRawData struct {
	ServiceCount float64 `bson:"service_count" json:"service_count"`
	HcodeCount   float64 `bson:"hcode_count" json:"hcode_count"`
	DmExpense    float64 `bson:"dm_expense" json:"dm_expense"`
	HtExpense    float64 `bson:"ht_expense" json:"ht_expense"`
	CopdExpense  float64 `bson:"copd_expense" json:"copd_expense"`
	CaExpense    float64 `bson:"ca_expense" json:"ca_expense"`
	PsyExpense   float64 `bson:"psy_expense" json:"psy_expense"`
	HdCvdExpense float64 `bson:"hd_cvd_expense" json:"hd_cvd_expense"`
}

type CidCountData struct {
	CidCount float64 `bson:"cid_count" json:"cid_count"`
}
