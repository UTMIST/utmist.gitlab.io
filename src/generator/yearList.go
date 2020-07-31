package generator

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const yearListSubstitution = "[//]: # years"

func getYearListString(name string, firstYear, lastYear, currentYear int) string {
	name = helpers.StringToSimplePath(name)
	yearListStr := "### "
	for y := lastYear; y >= firstYear; y-- {
		filepath := fmt.Sprintf("%s%s", helpers.ContentDirectory, name)
		if y != lastYear {
			filepath = fmt.Sprintf("%s-%d", filepath, y)
		}

		filepath = fmt.Sprintf("%s%s", filepath, helpers.MarkdownExt)
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		yearStr := fmt.Sprintf("%d-%d", y, y+1)
		if y == currentYear {
			yearStr = fmt.Sprintf("**%s**", yearStr)
		} else if y == lastYear {
			yearStr = fmt.Sprintf("[%s](../%s)", yearStr, name)
		} else {
			yearStr = fmt.Sprintf("[%s](../%s-%d)", yearStr, name, y)
		}
		yearListStr = fmt.Sprintf("%s%s | ", yearListStr, yearStr)
	}

	return yearListStr[:len(yearListStr)-3]
}
