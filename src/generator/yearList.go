package generator

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const yearListSubstitution = "[//]: # years"

func getYearListString(name string, firstYear, lastYear, currentYear int) string {
	name = helpers.StringToSimplePath(name)
	yearListStr := "### "
	for y := lastYear; y >= firstYear; y-- {
		yearStr := fmt.Sprintf("%d-%d", y, y+1)
		if y == currentYear {
			yearStr = fmt.Sprintf("**%s**", yearStr)
		} else if y == lastYear {
			yearStr = fmt.Sprintf("[%s](../%s)", yearStr, name)
		} else {
			yearStr = fmt.Sprintf("[%s](../%s-%d)", yearStr, name, y)
		}
		yearListStr = fmt.Sprintf("%s%s", yearListStr, yearStr)
		if y != firstYear {
			yearListStr = fmt.Sprintf("%s%s", yearListStr, " | ")
		}
	}

	return yearListStr
}
