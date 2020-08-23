package position

import (
	"fmt"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const teamPathPrefix = "team/"

// MakeList creates a list of open position lines.
func MakeList(positions *[]Position, deptPage bool, posType string) []string {

	var lines []string

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
			deptStr = fmt.Sprintf(
				", [%s](/%s%s)",
				pos.Department,
				teamPathPrefix,
				helpers.StringToSimplePath(pos.Department))
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
