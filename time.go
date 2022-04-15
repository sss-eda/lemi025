package lemi025

import (
	"fmt"
	"time"
)

// Time TODO
type Time struct {
	Year   uint8
	Month  uint8
	Day    uint8
	Hour   uint8
	Minute uint8
	Second uint8
}

// NewTime TODO
func NewTime(
	year uint8,
	month uint8,
	day uint8,
	hour uint8,
	minute uint8,
	second uint8,
) (*Time, error) {
	_, err := time.Parse(
		"2006-01-02 15:04:05",
		fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second),
	)
	if err != nil {
		return nil, err
	}

	return &Time{year, month, day, hour, minute, second}, nil
}
