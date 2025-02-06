package utils

import "fmt"

type FullMonthData struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type PatientData struct {
	Year    int     `bson:"year" json:"year"`
	Month   int     `bson:"month" json:"month"`
	Patient float64 `bson:"patient" json:"patient"`
}
type ExpenseData struct {
	Year    int     `bson:"year" json:"year"`
	Month   int     `bson:"month" json:"month"`
	Expense float64 `bson:"expense" json:"expe"`
}

func GenerateFullMonths(year int) []FullMonthData {
	fullMonths := make([]FullMonthData, 12)

	for i := 0; i < 12; i++ {
		currentYear := year
		if i >= 9 {
			currentYear++
		}
		month := (i + 4) % 12
		if month == 0 {
			month = 12
		}
		fullMonths[i] = FullMonthData{Year: currentYear, Month: month}
	}

	return fullMonths
}

func FillPatientResults(fullMonths []FullMonthData, data []PatientData) []float64 {
	result := make([]float64, len(fullMonths))

	dataMap := make(map[string]float64)
	for _, d := range data {
		key := fmt.Sprintf("%d-%d", d.Year, d.Month)
		dataMap[key] = d.Patient
	}

	for i, item := range fullMonths {
		key := fmt.Sprintf("%d-%d", item.Year, item.Month)
		if val, exists := dataMap[key]; exists {
			result[i] = val
		}
	}

	return result
}

func FillExpenseResults(fullMonths []FullMonthData, data []ExpenseData) []float64 {
	result := make([]float64, len(fullMonths))

	dataMap := make(map[string]float64)
	for _, d := range data {
		key := fmt.Sprintf("%d-%d", d.Year, d.Month)
		dataMap[key] = RoundToTwoDecimalPlaces(d.Expense)
	}

	for i, item := range fullMonths {
		key := fmt.Sprintf("%d-%d", item.Year, item.Month)
		if val, exists := dataMap[key]; exists {
			result[i] = val
		}
	}

	return result
}
