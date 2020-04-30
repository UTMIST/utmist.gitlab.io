package department

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
)

const deptListStart = "### **Departments & Groups**"
const execListStart = "### **Leadership**"
const joinParaStart = "### **Joining Us**"
const teamFilename = "content/team.md"

// GenerateDeptList generates a list of departments.
func GenerateDeptList(lines *[]string) {

	newLines := []string{}
	// Get list of departments and generate a line for each.
	for _, dept := range helpers.GetDeptNames(false) {
		deptLine := fmt.Sprintf("- [%s](%s)",
			dept, helpers.StringToFileName(dept))
		newLines = append(newLines, deptLine)
	}

	(*lines) = append(*lines, newLines...)
}

// GenerateTeamPage generates a page for the UTMIST team and open positions.
func GenerateTeamPage(
	associates *[]associate.Associate,
	positions *[]position.Position,
	descriptions *map[string]string) {

	// Start with the header and list of departments.
	lines := append(helpers.GenerateHeader("Our Team", "0001-01-01"),
		(*descriptions)["Team"], "", helpers.Breakline, "", deptListStart)
	GenerateDeptList(&lines)

	// Insert lists of execs.
	lines = append(lines, execListStart)
	associate.GenerateExecList(&lines, associates,
		(*descriptions)["Leadership"])

	// Add join prompt paragraphy and insert discord link.
	lines = append(lines, joinParaStart, (*descriptions)["Joining"])
	helpers.InsertDiscordLink(&lines)
	helpers.OverwriteWithLines(teamFilename, lines)
}
