package associate

import (
	"fmt"
	"strconv"
	"strings"
)

const associatesSheetRange = 17

const websiteLink = 0
const linkedinLink = 1
const gitlabLink = 2
const gitlhubLink = 3
const facebookLink = 4
const twitterLink = 5

const acd = "Academics"
const com = "Communications"
const ext = "External"
const fin = "Finance"
const lgs = "Logistics"
const mkt = "Marketing"
const osg = "Oversight"

// Associate represents an associateutive member's database row.
type Associate struct {
	Email       string
	PhoneNumber string

	FirstName     string
	PreferredName string
	LastName      string

	Discipline     string
	ProfilePicture string

	Departments []string
	Position    string
	VP          bool

	Website  string
	LinkedIn string
	GitLab   string
	GitHub   string
	Facebook string
	Twitter  string

	Retired int
}

// Return selected social media link for user.
func (e *Associate) getLink(link int) string {
	switch link {
	case websiteLink:
		if len(e.Website) > 0 {
			return e.Website
		}
	case linkedinLink:
		if len(e.LinkedIn) > 0 {
			return fmt.Sprintf("https://linkedin.com/in/%s/", e.LinkedIn)
		}
	case gitlabLink:
		if len(e.GitLab) > 0 {
			return fmt.Sprintf("https://www.gitlab.com/%s/", e.GitLab)
		}
	case gitlhubLink:
		if len(e.GitHub) > 0 {
			return fmt.Sprintf("https://www.github.com/%s/", e.GitHub)
		}
	case facebookLink:
		if len(e.Facebook) > 0 {
			return fmt.Sprintf("https://www.facebook.com/%s/", e.Facebook)
		}
	case twitterLink:
		if len(e.Twitter) > 0 {
			return fmt.Sprintf("https://www.twitter.com/%s/", e.Twitter)
		}
	}

	return ""
}

// Get list of departments.
func getDepartments() []string {
	return []string{
		acd, com, ext, fin, lgs, mkt, osg,
	}
}

// LoadAssociate loads an associate from a spreadsheet row.
func LoadAssociate(data []interface{}) Associate {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < associatesSheetRange; i++ {
		data = append(data, "")
	}

	associate := Associate{
		Email:         data[0].(string),
		FirstName:     data[1].(string),
		PreferredName: data[2].(string),
		LastName:      data[3].(string),
		PhoneNumber:   data[4].(string),
		Position:      data[5].(string),
		Departments:   strings.Split(data[6].(string), ", "),
		Discipline:    data[8].(string),
		Website:       data[10].(string),
		Facebook:      data[11].(string),
		Twitter:       data[12].(string),
		LinkedIn:      data[13].(string),
		GitHub:        data[14].(string),
		GitLab:        data[15].(string),
		Retired: func(yearObj interface{}) int {
			retired, err := strconv.Atoi(yearObj.(string))
			if err != nil {
				retired = -1
			}
			return retired
		}(data[16]),
	}
	return associate
}
