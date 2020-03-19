package associate

import (
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const associatesSheetRange = 17

const websiteLink = 0
const linkedinLink = 1
const gitlabLink = 2
const gitlhubLink = 3
const facebookLink = 4
const twitterLink = 5

// LoadAssociate loads an associate from a spreadsheet row.
func LoadAssociate(data []interface{}) []Associate {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < associatesSheetRange; i++ {
		data = append(data, "")
	}

	// Create the base associate.
	associate := Associate{
		Email:         data[0].(string),
		FirstName:     data[1].(string),
		PreferredName: data[2].(string),
		LastName:      data[3].(string),
		PhoneNumber:   data[4].(string),
		Position:      data[5].(string),
		Department:    data[6].(string),
		Discipline:    data[8].(string),
		Website:       data[10].(string),
		Facebook:      data[11].(string),
		Twitter:       data[12].(string),
		LinkedIn:      data[13].(string),
		GitHub:        data[14].(string),
		GitLab:        data[15].(string),
		Graduated:     helpers.InterfaceToYear(data[16]),
	}

	// Add a single entry for any alumni.
	if associate.hasGraduated() {
		associate.Department = "Alumni"
		return []Associate{associate}
	}

	// Create a version for associate for every department-position pair.
	entries := []Associate{}
	positions := strings.Split(data[5].(string), ",")
	departments := strings.Split(data[6].(string), ",")
	count := len(positions)
	if len(departments) > count {
		count = len(departments)
	}
	for i := 0; i < count; i++ {
		associate.Department = strings.Trim(departments[i], " ")
		associate.Position = strings.Trim(positions[i], " ")
		entries = append(entries, associate)
	}

	return entries
}
