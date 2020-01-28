package generator

import "fmt"

// Exec represents an executive member's database row.
type Exec struct {
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
	GitHub   string
	GitLab   string
	Facebook string
	Twitter  string

	Retired int
}

func (e *Exec) getLink(link int) string {
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

func getDepartments() []string {
	return []string{
		acd, com, ext, fin, lgs, mkt, osg,
	}
}

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
