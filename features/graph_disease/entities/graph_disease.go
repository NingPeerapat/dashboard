package entities

type DiseaseRequest struct {
	Year     int    `json:"year"`
	Area     string `json:"area"`
	Province string `json:"province"`
	District string `json:"district"`
	Hcode    string `json:"hcode"`
}

type DiseaseResponse struct {
	Status  bool          `json:"status"`
	Message string        `json:"message"`
	Result  []DiseaseData `json:"result"`
}

type DiseaseData struct {
	DiseaseName string    `json:"diseaseName"`
	Data        []float64 `json:"data"`
}

type DiseaseExpenseData struct {
	Year         int     `bson:"year" json:"year"`
	Month        int     `bson:"month" json:"month"`
	DMExpense    float64 `bson:"dm_expense" json:"dm_expense"`
	HtExpense    float64 `bson:"ht_expense" json:"ht_expense"`
	CopdExpense  float64 `bson:"copd_expense" json:"copd_expense"`
	CaExpense    float64 `bson:"ca_expense" json:"ca_expense"`
	PsyExpense   float64 `bson:"psy_expense" json:"psy_expense"`
	HdCvdExpense float64 `bson:"hd_cvd_expense" json:"hd_cvd_expense"`
}

type DiseasePatientData struct {
	Year          int     `bson:"year" json:"year"`
	Month         int     `bson:"month" json:"month"`
	DMCidCount    float64 `bson:"dm_cid_count" json:"dm_cid_count"`
	HtCidCount    float64 `bson:"ht_cid_count" json:"ht_cid_count"`
	CopdCidCount  float64 `bson:"copd_cid_count" json:"copd_cid_count"`
	CaCidCount    float64 `bson:"ca_cid_count" json:"ca_cid_count"`
	PsyCidCount   float64 `bson:"psy_cid_count" json:"psy_cid_count"`
	HdCvdCidCount float64 `bson:"hd_cvd_cid_count" json:"hd_cvd_cid_count"`
}
