package dto

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
