package position

import (
	"fmt"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const positionBaseCopyPath = "./assets/recruitment_preface.md"
const positionPageCopyPath = "./assets/recruitment.md"
const positionPagePath = "./content/recruitment.md"

// MakeList creates a list of open position lines.
func MakeList(positions *[]Position, deptPage bool, posType string) []string {

	var lines []string
	if deptPage {
		lines = helpers.ReadContentLines(positionBaseCopyPath)
		lines = append(
			[]string{helpers.Breakline, "### **Open Positions**"},
			lines...)
	} else {
		lines = []string{"",
			helpers.Breakline, "",
			fmt.Sprintf("## **%s Positions**", posType)}
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
		head := fmt.Sprintf("##### **%s**%s", pos.Title, func() string {
			if deptPage {
				return ""
			}
			return fmt.Sprintf(", [%s](%s)",
				pos.Department,
				helpers.StringToFileName(pos.Department))
		}())

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

// GeneratePage generates a page for recruitment.
func GeneratePage(positions *[]Position) {
	execPositions := []Position{}
	assocPositions := []Position{}

	for _, pos := range *positions {
		if pos.IsExec() {
			execPositions = append(execPositions, pos)
			continue
		}
		assocPositions = append(assocPositions, pos)
	}

	lines := helpers.ReadFileBase(positionPageCopyPath, len(*positions), 6)
	lines = append(lines, helpers.GetJoinLines()...)
	lines = append(lines, MakeList(&execPositions, false, "Executive")...)
	lines = append(lines, MakeList(&assocPositions, false, "Associate")...)

	helpers.OverwriteWithLines(positionPagePath, lines)
}
