package associate

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// MakeDepartmentList generates a string list of departments
func MakeDepartmentList(depts *[]string) []string {
	list := []string{}
	for _, dept := range *depts {
		list = append(
			list,
			fmt.Sprintf(
				"- [%s](../%s)",
				dept,
				helpers.StringToSimplePath(dept)))
	}
	return list
}
