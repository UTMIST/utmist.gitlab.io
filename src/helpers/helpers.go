package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// FormatDateEST formats a date from EST.
func FormatDateEST(dateStr string) time.Time {

	if len(dateStr) == 0 {
		return time.Now().AddDate(1, 0, 0)
	}

	layout := parseDateLayout
	if strings.Count(dateStr, " ") == 1 {
		layout = parseDateTimeLayout
	}

	// Load location and parse the time in that location.
	toronto, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}
	}
	dateTime, err := time.ParseInLocation(layout, dateStr, toronto)
	if err != nil {
		return time.Time{}
	}

	return dateTime
}

// GetDeptNames returns a list of department names.
func GetDeptNames(year int) []string {
	depts := strings.Split(os.Getenv(fmt.Sprintf("DEPTS_%d", year)), ",")
	for i := 0; i < len(depts); i++ {
		depts[i] = strings.TrimSpace(depts[i])
	}

	return depts
}

// GetURLBase shortens a URL to a short preview.
func GetURLBase(link string) string {
	if index := strings.Index(link, "//"); index >= 0 {
		link = link[index+2:]
	}
	if index := strings.Index(link, "/"); index >= 0 {
		link = link[:index]
	}
	return link
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

// ColonRemainder returns the remainder of a string after the first colon.
func ColonRemainder(line string) string {
	remainder := strings.TrimSpace(line[strings.Index(line, ":")+1:])
	return strings.Trim(remainder, "\"")
}
