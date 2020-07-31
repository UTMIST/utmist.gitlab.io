package associate

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// MakeDepartmentList generates a string list of departments
func MakeDepartmentList(depts *[]string, lastYear, currentYear int) []string {
	list := []string{}
	for _, dept := range *depts {

		filename := fmt.Sprintf("%s", helpers.StringToSimplePath(dept))
		if lastYear != currentYear {
			filename = fmt.Sprintf("%s-%d", filename, currentYear)
		}

		filepath := fmt.Sprintf(
			"%s%s%s",
			helpers.ContentDirectory,
			filename,
			helpers.MarkdownExt)
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		list = append(list, fmt.Sprintf("- [%s](../%s)", dept, filename))
	}
	return list
}
