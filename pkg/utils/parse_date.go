package utils

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date: %v", err)
	}
	return date, nil
}
