package helpers

import (
	"os"
	"strconv"
	"strings"
)

// GetYearRange returns the first and last years of range from env.
func GetYearRange() (int, int) {
	years := strings.Split(os.Getenv("YEARS"), "_")
	firstYear, err := strconv.Atoi(years[0])
	if err != nil {
		panic(err)
	}
	lastYear, err := strconv.Atoi(years[1])
	if err != nil {
		panic(err)
	}

	return firstYear, lastYear
}
