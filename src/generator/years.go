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
		thisYearStr := fmt.Sprintf("%d-%d", y, y+1)
		if y == currentYear {
			yearListStr = fmt.Sprintf("%s**%s** | ", yearListStr, thisYearStr)
			continue
		}

		filepathParts := strings.Split(filepath, "/")
		if y == lastYear {
			filepathParts = append(filepathParts[:2], filepathParts[3:]...)
		} else if currentYear == lastYear {
			filepathParts = append(filepathParts[:2], append(
				[]string{thisYearStr},
				filepathParts[2:]...)...)
		} else {
			filepathParts[2] = thisYearStr
		}

		if _, err := os.Stat(strings.Join(filepathParts, "/")); err != nil {
			continue
		}

		yearListStr = fmt.Sprintf(
			"%s[%s](/%s) | ",
			yearListStr,
			thisYearStr,
			strings.Join(filepathParts[1:len(filepathParts)-1], "/"))
	}

	return yearListStr[:len(yearListStr)-3]
}
