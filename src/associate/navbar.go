package associate

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// Where the department list starts in the config.
const start = "        - title: Our Team"

// GenerateNavbarDeptLinks generates event links for the navbar dropdown menu.
func GenerateNavbarDeptLinks(lines *[]string) {
	depts := GetDepartmentNames()
	helpers.StitchPageLink(lines, depts[:len(depts)-1], "/team/", start)
}
