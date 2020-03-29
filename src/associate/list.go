package associate

import (
	"fmt"
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
)

const deptListStart = "## **Departments**"
const execListStart = "## **Leadership**"

const teamCopyFilename = "assets/team.md"
const teamFilename = "content/team.md"

// GenerateExecList generates a list of executive members.
func GenerateExecList(lines *[]string, associates *[]Associate) {

	// Add each exec to the list.
	execs := []Associate{}
	for _, associate := range *associates {
		if associate.isExec() && !associate.hasGraduated() {
			execs = append(execs, associate)
		}
	}
	sort.Sort(List(execs))

	newLines := []string{}
	for _, exec := range execs {
		execLine := exec.getLine("", false, true)
		newLines = append(newLines, execLine)
	}
	newLines = append(newLines, "")
	helpers.StitchIntoLines(lines, &newLines, execListStart, 1)
}

// GenerateDeptList generates a list of departments.
func GenerateDeptList(lines *[]string) {

	newLines := []string{}
	// Get list of departments and generate a line for each.
	for _, dept := range GetDepartmentNames(false) {
		deptLine := fmt.Sprintf("- [%s](%s)", dept, helpers.StringToFileName(dept))
		newLines = append(newLines, deptLine)
	}

	helpers.StitchIntoLines(lines, &newLines, deptListStart, 1)

}

// GenerateTeamPage generates a page for the UTMIST team and open positions.
func GenerateTeamPage(associates *[]Associate, positions *[]position.Position) {
	lines := helpers.ReadContentLines(teamCopyFilename)
	GenerateDeptList(&lines)
	GenerateExecList(&lines, associates)
	lines = append(lines, position.MakeList(positions, false)...)

	helpers.InsertDiscordLink(&lines)
	helpers.OverwriteWithLines(teamFilename, lines)
}
