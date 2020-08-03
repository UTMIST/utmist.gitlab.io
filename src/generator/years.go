package generator

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const yearListSubstitution = "[//]: # years"

func getYearListString(filepath string, currentYear int) string {
	yearListStr := "### "
	firstYear, lastYear, err := helpers.GetYearRange(os.Getenv("YEARS"))
	if err != nil {
		panic(err)
	}

	filepathParts := strings.Split(filepath, "/")
	latestVersion := false
	if _, _, err := helpers.GetYearRange(filepathParts[2]); err != nil {
		latestVersion = true
	}
	for y := lastYear; y >= firstYear; y-- {
		yearStr := fmt.Sprintf("%d-%d", y, y+1)

		thisFilepathParts := strings.Split(filepath, "/")[:]

		if latestVersion && y != lastYear {
			thisFilepathParts = append(thisFilepathParts[:2], append([]string{yearStr}, thisFilepathParts[2:]...)...)
		} else {
			thisFilepathParts[2] = yearStr
		}

		thisFilepath := strings.Join(thisFilepathParts, "/")
		if _, err := os.Stat(thisFilepath); err != nil {
			continue
		}

		if y == currentYear {
			yearStr = fmt.Sprintf("**%s**", yearStr)
		} else if y == lastYear {
			yearStr = fmt.Sprintf("[%s](../%s)", yearStr, filepath)
		} else {
			yearStr = fmt.Sprintf("[%s](../%s-%d)", yearStr, filepath, y)
		}
		yearListStr = fmt.Sprintf("%s%s | ", yearListStr, yearStr)
	}

	return yearListStr[:len(yearListStr)-3]
}
