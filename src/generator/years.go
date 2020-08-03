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

	for y := lastYear; y >= firstYear; y-- {
		currentYearStr := fmt.Sprintf("%d-%d", currentYear, currentYear+1)
		thisYearStr := fmt.Sprintf("%d-%d", y, y+1)

		newPath := strings.ReplaceAll(filepath, currentYearStr, thisYearStr)
		if _, err := os.Stat(newPath); err != nil {
			continue
		}

		var yearLink string
		if y == currentYear {
			yearLink = fmt.Sprintf("**%s**", thisYearStr)
		} else {
			yearLink = fmt.Sprintf("[%s](/%s)",
				thisYearStr,
				strings.TrimSuffix(newPath[strings.Index(newPath, "/")+1:],
					helpers.PageIndex))
		}
		yearListStr = fmt.Sprintf("%s%s | ", yearListStr, yearLink)
	}

	return yearListStr[:len(yearListStr)-3]
}
