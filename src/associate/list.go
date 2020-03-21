package associate

import (
	"fmt"
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
)

const teamFileBase = "./assets/team.md"
const deptListStart = "## **Departments**"
const execListStart = "## **Leadership**"
const orgListStart = "| Departments |     | Leadership |"

const teamCopyFilename = "assets/team.md"
const teamFilename = "content/team.md"

// GenerateOrgList generates a table if columns of departments and execs.
func GenerateOrgList(lines *[]string, associates *[]Associate) {
	// Get list of departments.
	depts := GetDepartmentNames(false)

	// Add each exec to the list.
	execs := []Associate{}
	for _, associate := range *associates {
		if associate.isExec() && !associate.hasGraduated() {
			execs = append(execs, associate)
		}
	}
	sort.Sort(List(execs))

	newLines := []string{}
	for index := 0; index < len(depts) || index < len(execs); index++ {
		line := fmt.Sprintf("|%s|%s|%s|",
			func() string {
				if index >= len(depts) {
					return ""
				}
				return fmt.Sprintf("[%s](%s)",
					depts[index],
					helpers.StringToFileName(depts[index]))
			}(),
			helpers.TablePad(12),
			func() string {
				if index >= len(execs) {
					return ""
				}
				return execs[index].getLine("", false, false)
			}())
		newLines = append(newLines, line)
	}
	newLines = append(newLines, "")
	helpers.StitchIntoLines(lines, &newLines, orgListStart, 1)
}

// GenerateTeamPage generates a page for the UTMIST team and open positions.
func GenerateTeamPage(associates *[]Associate, positions *[]position.Position) {
	lines := helpers.ReadContentLines(teamCopyFilename)
	GenerateOrgList(&lines, associates)
	lines = append(lines, position.MakeList(positions, false)...)

	helpers.InsertDiscordLink(&lines)
	helpers.OverwriteWithLines(teamFilename, lines)
}
