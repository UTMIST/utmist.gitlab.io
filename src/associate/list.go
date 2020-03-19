package associate

import (
	"fmt"
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const teamFileBase = "./assets/team.md"
const tablePadder = "   "
const deptListStart = "## **Departments**"
const execListStart = "## **Leadership**"

// GenerateDeptList generates the list of departments into team.md.
func GenerateDeptList(lines *[]string) {
	// Add each dept to the list.
	newLines := []string{}
	for _, dept := range GetDepartmentNames() {
		if dept == alm {
			continue
		}
		line := fmt.Sprintf("- [%s](../%s)", dept, helpers.StringToFileName(dept))
		newLines = append(newLines, line)
	}
	newLines = append(newLines, "")
	helpers.StitchIntoLines(lines, &newLines, deptListStart, 1)
}

// GenerateVPList generates a list of VPs into team.md.
func GenerateVPList(lines *[]string, associates []Associate) {
	// Add each dept to the list.
	execs := []Associate{}
	for _, associate := range associates {
		if associate.isExec() && !associate.hasGraduated() {
			execs = append(execs, associate)
		}
	}
	sort.Sort(List(execs))

	newLines := []string{}
	for _, exec := range execs {
		newLines = append(newLines, exec.getLine("", false))
	}

	newLines = append(newLines, "")
	helpers.StitchIntoLines(lines, &newLines, execListStart, 1)
}
