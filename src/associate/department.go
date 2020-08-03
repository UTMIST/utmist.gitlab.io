package associate

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// MakeDepartmentList generates a string list of departments
func MakeDepartmentList(depts *[]string, year int) []string {
	list := []string{}
	for _, dept := range *depts {

		filename := helpers.StringToSimplePath(dept)
		filepath := fmt.Sprintf(
			"%steam/%s/%s/%s",
			helpers.ContentDirectory,
			fmt.Sprintf("%d-%d", year, year+1),
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
