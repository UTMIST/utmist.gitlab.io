package position

import (
	"fmt"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const positionPagePath = "./content/recruitment.md"

// MakeList creates a list of open position lines.
func MakeList(positions *[]Position, deptPage bool,
	posType, desc string) []string {

	var lines []string
	if deptPage {
		lines = []string{"", helpers.Breakline, helpers.OpenPositions, desc}
	} else {
		lines = []string{"", helpers.Breakline,
			fmt.Sprintf("### **%s Positions**", posType)}
	}

	if positions == nil || len(*positions) == 0 {
		if deptPage {
			return []string{}
		}
		return append(lines,
			fmt.Sprintf(
				"We currently don't have %s openings. Check back later!",
				posType))
	}

	for _, pos := range *positions {
		deptStr := ""
		if !deptPage {
			deptStr = fmt.Sprintf(", [%s](/team/%s)", pos.Department,
				helpers.StringToFileName(pos.Department))
		}
		head := fmt.Sprintf("##### **%s**%s", pos.Title, deptStr)

		lines = append(lines, head)
		if len(pos.Description) > 0 {
			lines = append(lines, pos.Description)
		}
		if len(pos.Requirements) > 0 {
			lines = append(lines,
				fmt.Sprintf("- _Requirements_: %s", pos.Requirements))
		}
		if len(pos.Instructions) > 0 {
			lines = append(lines,
				fmt.Sprintf("- _Instructions_: %s", pos.Instructions))
		}
		if pos.Deadline != time.Unix(0, 0) {
			lines = append(lines,
				fmt.Sprintf("- _Application Due_: **%s**",
					pos.Deadline.Format(helpers.PrintDateLayout)))
		}
		lines = append(lines, "")
	}

	return lines
}

// GenerateList generates a page for recruitment.
func GenerateList(positions *[]Position, descriptions *map[string]string) {
	execPositions := []Position{}
	assocPositions := []Position{}

	for _, pos := range *positions {
		if pos.IsExec() {
			execPositions = append(execPositions, pos)
			continue
		}
		assocPositions = append(assocPositions, pos)
	}

	lines := append(helpers.GenerateHeader("Join Us", "0001-01-03"),
		helpers.GetJoinLines((*descriptions)["Joining"])...)
	lines = append(lines, MakeList(
		&execPositions, false,
		"Executive", (*descriptions)["Recruitment"])...)
	lines = append(lines, MakeList(
		&assocPositions, false,
		"Associate", (*descriptions)["Recruitment"])...)

	helpers.OverwriteWithLines(positionPagePath, lines)
}
