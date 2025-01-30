package entities

type CaRequest struct {
	Year     int    `json:"year"`
	Area     string `json:"area"`
	Province string `json:"province"`
	District string `json:"district"`
	Hcode    string `json:"hcode"`
}

type CaResponse struct {
	Status  bool     `json:"status"`
	Message string   `json:"message"`
	Result  []CaData `json:"result"`
}

type CaData struct {
	DiseaseName string    `json:"diseaseName"`
	Data        []float64 `json:"data"`
}

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

type CaPatientData struct {
	Year                 int     `bson:"year" json:"year"`
	Month                int     `bson:"month" json:"month"`
	CaCidCount           float64 `bson:"ca_cid_count" json:"ca_cid_count"`
	LungCaCidCount       float64 `bson:"lung_ca_cid_count" json:"lung_ca_cid_count"`
	BreastCaCidCount     float64 `bson:"breast_ca_cid_count" json:"breast_ca_cid_count"`
	CervicalCaCidCount   float64 `bson:"cervical_ca_cid_count" json:"cervical_ca_cid_count"`
	LiverCaCidCount      float64 `bson:"liver_ca_cid_count" json:"liver_ca_cid_count"`
	ColorectalCaCidCount float64 `bson:"colorectal_ca_cid_count" json:"colorectal_ca_cid_count"`
}
