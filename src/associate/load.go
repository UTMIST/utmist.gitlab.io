package associate

import (
	"strings"
)

const associateRowLength = 14
const entryRowLength = 4

// LoadAssociate loads an associate from a spreadsheet row.
func LoadAssociate(data []interface{}) Associate {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < associateRowLength; i++ {
		data = append(data, "")
	}

	return Associate{
		data[0].(string),
		data[1].(string),
		data[2].(string),
		data[3].(string),
		data[4].(string),
		data[5].(string),
		data[6].(string),
		data[7].(string),
		data[8].(string),
		data[9].(string),
		data[10].(string),
		data[11].(string),
		data[12].(string),
		data[13].(string),
	}
}

// LoadEntries loads an associate entry from a spreadsheet row.
func LoadEntries(data []interface{}, associates *map[string]Associate) []Entry {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < entryRowLength; i++ {
		data = append(data, "")
	}
	positions := strings.Split(data[4].(string), ",")
	departments := strings.Split(data[5].(string), ",")

	entries := []Entry{}

	for i := 0; i < len(positions); i++ {
		if i >= len(departments) {
			break
		}

		email := data[3].(string)
		associate := (*associates)[email]

		entries = append(
			entries,
			Entry{
				email,
				strings.TrimSpace(positions[i]),
				strings.TrimSpace(departments[i]),
				&associate})
	}

	return entries
}
