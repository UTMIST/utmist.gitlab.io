package position

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const positionSheetRange = 6

// Load loads a position from a spreadsheet row.
func Load(data []interface{}) Position {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < positionSheetRange; i++ {
		data = append(data, "")
	}

	position := Position{
		Title:        data[0].(string),
		Department:   data[1].(string),
		Description:  data[2].(string),
		Requirements: data[3].(string),
		Instructions: data[4].(string),
	}

	if len(data[5].(string)) > 0 {
		position.Deadline = helpers.FormatDateEST(data[5].(string))
	}

	return position
}
