package associate

import (
	"fmt"
	"log"
	"os"
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// MakeDepartmentList generates a string list of departments
func MakeDepartmentList(depts *[]string, year int) []string {
	list := []string{}

	_, lastYear, err := helpers.GetYearRange(os.Getenv("YEARS"))
	if err != nil {
		return list
	}

	sort.Strings(*depts)
	for _, dept := range *depts {

		yearStr := fmt.Sprintf("%d-%d/", year, year+1)
		if year == lastYear {
			yearStr = ""
		}

		filename := helpers.StringToSimplePath(dept)
		filepath := fmt.Sprintf(
			"%steam/%s%s/%s",
			helpers.ContentDirectory,
			yearStr,
			helpers.StringToSimplePath(dept),
			helpers.PageIndex)
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		list = append(list, fmt.Sprintf("- [%s](../%s)", dept, filename))
	}

	return list
}
