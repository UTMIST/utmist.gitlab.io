package position

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// MakeList creates a list of open position lines.
func MakeList(positions *[]Position, deptPage bool) []string {

	if positions == nil || len(*positions) == 0 {
		return []string{}
	}

	lines := helpers.ReadFileBase(positionBasePath, len(*positions), 5)

	for _, pos := range *positions {
		head := fmt.Sprintf("##### **%s**%s", pos.Title, func() string {
			if deptPage {
				return ""
			}
			return fmt.Sprintf(", [%s](%s)",
				pos.Department,
				helpers.StringToFileName(pos.Department))
		}())

		desc := fmt.Sprintf("- _Description_: %s", pos.Description)
		reqs := fmt.Sprintf("- _Requirements_: %s", pos.Requirements)
		inst := fmt.Sprintf("- _Instructions_: %s", pos.Instructions)

		lines = append(lines, []string{head, desc, reqs, inst, "\n"}...)
	}

	return lines
}
