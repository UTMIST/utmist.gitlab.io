package position

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const recruitSheetRange = 5

const recruitBasePath = "./assets/positions.md"

// Position represents an open associate position at UTMIST.
type Position struct {
	Title        string
	Department   string
	Description  string
	Requirements string
	Instructions string
}

// LoadPosition loads a position from a spreadsheet row.
func LoadPosition(data []interface{}) Position {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < recruitSheetRange; i++ {
		data = append(data, "")
	}

	position := Position{
		Title:        data[0].(string),
		Department:   data[1].(string),
		Description:  data[2].(string),
		Requirements: data[3].(string),
		Instructions: data[4].(string),
	}

	return position
}

// MakeTable creates a list of open position lines.
func MakeTable(positions []Position) []string {

	if len(positions) == 0 {
		return []string{}
	}

	lines := helpers.ReadFileBase(recruitBasePath, len(positions), 6)

	for _, pos := range positions {
		posListing := fmt.Sprintf("|%s|%s|[%s](%s)|%s|%s|%s|%s|%s|%s|",
			pos.Title,
			helpers.TablePadder,
			pos.Department,
			helpers.StringToFileName(pos.Department),
			helpers.TablePadder,
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
