package department

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
)

const deptListStart = "## **Departments**"
const teamCopyFilename = "assets/team.md"
const teamFilename = "content/team.md"

// GenerateDeptList generates a list of departments.
func GenerateDeptList(lines *[]string) {

	newLines := []string{}
	// Get list of departments and generate a line for each.
	for _, dept := range GetDeptNames(false) {
		deptLine := fmt.Sprintf("- [%s](%s)", dept, helpers.StringToFileName(dept))
		newLines = append(newLines, deptLine)
	}

	helpers.StitchIntoLines(lines, &newLines, deptListStart, 1)

}

// GenerateTeamPage generates a page for the UTMIST team and open positions.
func GenerateTeamPage(associates *[]associate.Associate, positions *[]position.Position) {
	lines := helpers.ReadContentLines(teamCopyFilename)
	GenerateDeptList(&lines)
	associate.GenerateExecList(&lines, associates)
	lines = append(lines, position.MakeList(positions, false)...)

	helpers.InsertDiscordLink(&lines)
	helpers.OverwriteWithLines(teamFilename, lines)
}
