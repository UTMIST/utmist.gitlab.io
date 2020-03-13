package helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const parseDateLayout = "01/02/2006 15:04:05"

// FormatDateEST formats a date from EST.
func FormatDateEST(dateStr string) time.Time {
	// UNDOCUMENTED.
	parts := strings.Split(dateStr, "/")
	for i := 0; i < 2; i++ {
		if len(parts[i]) == 1 {
			parts[i] = fmt.Sprintf("0%s", parts[i])
		}
	}
	dateStr = strings.Join(parts, "/")

	// UNDOCUMENTED.
	parts = strings.Split(dateStr, " ")
	if parts[1][1] == ':' {
		parts[1] = fmt.Sprintf("0%s", parts[1])
	}
	dateStr = strings.Join(parts, " ")

	// Load location and parse the time in that location.
	toronto, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}
	}
	dateTime, err := time.ParseInLocation(parseDateLayout, dateStr, toronto)
	if err != nil {
		return time.Time{}
	}

	return dateTime
}

// InterfaceToYear produces a year A.E. from an interface (usually a string).
func InterfaceToYear(yearObj interface{}) int {
	if year, err := strconv.Atoi(yearObj.(string)); err == nil {
		return year
	}
	return -1
}

// PadDateWithIndex pads a year with 0s and Jan 1st.
func PadDateWithIndex(index int) string {
	padded := fmt.Sprintf("%04d-01-01\n", index)
	return padded
}

// PositionNumber returns the number of positions open for a position from a string.
func PositionNumber(numStr string) int {
	if num, err := strconv.Atoi(numStr); err == nil {
		return num
	}
	return 1
}
