package generator

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

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
		_, _, err := helpers.GetYearRange(filepathParts[2])
		filepathHasYear := err == nil
		if y == lastYear {
			filepathParts = append(filepathParts[:2], filepathParts[3:]...)
		} else if currentYear == lastYear && !filepathHasYear {
			filepathParts = append(filepathParts[:2], append(
				[]string{thisYearStr},
				filepathParts[2:]...)...)
		} else {
			filepathParts[2] = thisYearStr
		}

		newFilepath := strings.Join(filepathParts, "/")
		_, err1 := os.Stat(newFilepath)

		fallbackFilepath := strings.Join(append(
			filepathParts[:2],
			append([]string{thisYearStr}, filepathParts[2:]...)...), "/")
		_, err2 := os.Stat(fallbackFilepath)

		if err1 != nil && err2 != nil {
			continue
		} else if err1 != nil {
			newFilepath = fallbackFilepath
		}

		filepathParts = strings.Split(newFilepath, "/")
		yearListStr = fmt.Sprintf(
			"%s[%s](/%s) | ",
			yearListStr,
			thisYearStr,
			strings.Join(filepathParts[1:len(filepathParts)-1], "/"))
	}

	return yearListStr[:len(yearListStr)-3]
}
