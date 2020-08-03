package helpers

import (
	"strconv"
	"strings"
)

// GetYearRange returns the first and last years of range from env.
func GetYearRange(yearStr string) (int, int, error) {
	years := strings.Split(yearStr, "-")
	firstYear, err := strconv.Atoi(years[0])
	if err != nil {
		return -1, -1, err
	}
	lastYear, err := strconv.Atoi(years[1])
	if err != nil {
		return -1, -1, err
	}

	return firstYear, lastYear, nil
}
