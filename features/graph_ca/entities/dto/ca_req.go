package dto

type CaRequest struct {
	Year     int    `json:"year"`
	Area     string `json:"area"`
	Province string `json:"province"`
	District string `json:"district"`
	Hcode    string `json:"hcode"`
}
