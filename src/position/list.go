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

	lines := helpers.ReadFileBase(func() string {
		if deptPage {
			return positionDeptBasePath
		}
		return positionBasePath
	}(), len(*positions), 6)

	for _, pos := range *positions {
		posListing := fmt.Sprintf("|%s|%s|%s%s|%s|%s|%s|%s|",
			pos.Title,
			helpers.TablePadder,
			func() string {
				if deptPage {
					return ""
				}
				return fmt.Sprintf("[%s](%s)|%s|",
					pos.Department,
					helpers.StringToFileName(pos.Department),
					helpers.TablePadder)
			}(),
			pos.Description,
			helpers.TablePadder,
			pos.Requirements,
			helpers.TablePadder,
			pos.Instructions,
		)
		lines = append(lines, posListing)
	}

	return lines
}
