package associate

import (
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const associatesSheetRange = 17

// Load loads an associate from a spreadsheet row.
func Load(data []interface{}) []Associate {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < associatesSheetRange; i++ {
		data = append(data, "")
	}

	// Create the base associate.
	assoc := Associate{
		FirstName:     data[0].(string),
		PreferredName: data[1].(string),
		LastName:      data[2].(string),

		UofTEmail:   data[3].(string),
		Email:       data[4].(string),
		PhoneNumber: data[5].(string),

		Position:   data[6].(string),
		Department: data[7].(string),
		Retirement: helpers.FormatDateEST(data[8].(string)),
		Discipline: data[9].(string),

		ProfilePicture: data[10].(string),
		Website:        data[11].(string),
		LinkedIn:       data[12].(string),
		GitHub:         data[13].(string),
		GitLab:         data[14].(string),
		Facebook:       data[15].(string),
		Twitter:        data[16].(string),
	}

	// Add a single entry for any alumni.
	if assoc.HasRetired() {
		assoc.Department = helpers.ALM
		return []Associate{assoc}
	}

	// Create a version for associate for every department-position pair.
	entries := []Associate{}
	positions := strings.Split(assoc.Position, ",")
	departments := strings.Split(assoc.Department, ",")
	count := len(positions)
	if len(departments) < count {
		count = len(departments)
	}
	for i := 0; i < count; i++ {
		assoc.Department = strings.Trim(departments[i], " ")
		assoc.Position = strings.Trim(positions[i], " ")
		entries = append(entries, assoc)
	}

	return entries
}
