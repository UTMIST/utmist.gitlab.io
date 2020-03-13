package associate

import (
	"fmt"
	"strings"
)

const adm = "Administration"
const acd = "Academics"
const ind = "Industry Relations"
const inf = "Infrastructure"
const mkt = "Marketing"
const prj = "Projects"
const twr = "Technical Writing"
const alm = "Alumni"

// GetDepartmentNames returns a list of department names.
func GetDepartmentNames() []string {
	return []string{
		adm, acd, ind, inf, mkt, prj, twr, alm,
	}
}

// Associate represents an associateutive member's database row.
type Associate struct {
	Email       string
	PhoneNumber string

	FirstName     string
	PreferredName string
	LastName      string

	Discipline     string
	ProfilePicture string

	Department string
	Position   string
	VP         bool

	Website  string
	LinkedIn string
	GitLab   string
	GitHub   string
	Facebook string
	Twitter  string

	Graduated int
}

// List defines a list of events.
type List []Associate

// Method Len() to implement sort.Sort.
func (a List) Len() int {
	return len(a)
}

// Method Less() to implement sort.Sort.
func (a List) Less(i, j int) bool {

	switch strings.Compare(a[i].Position, a[j].Position) {
	case -1:
		return true
	case 1:
		return false
	}

	switch strings.Compare(a[i].LastName, a[j].LastName) {
	case -1:
		return true
	case 1:
		return false
	}

	switch strings.Compare(a[i].FirstName, a[j].FirstName) {
	case -1:
		return true
	case 1:
		return false
	}

	return false
}

// Method Swap() to implement sort.Sort.
func (a List) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Return selected social media link for user.
func (a *Associate) getLink(link int) string {
	switch link {
	case websiteLink:
		if len(a.Website) > 0 {
			return a.Website
		}
	case linkedinLink:
		if len(a.LinkedIn) > 0 {
			return fmt.Sprintf("https://linkedin.com/in/%s/", a.LinkedIn)
		}
	case gitlabLink:
		if len(a.GitLab) > 0 {
			return fmt.Sprintf("https://www.gitlab.com/%s/", a.GitLab)
		}
	case gitlhubLink:
		if len(a.GitHub) > 0 {
			return fmt.Sprintf("https://www.github.com/%s/", a.GitHub)
		}
	case facebookLink:
		if len(a.Facebook) > 0 {
			return fmt.Sprintf("https://www.facebook.com/%s/", a.Facebook)
		}
	case twitterLink:
		if len(a.Twitter) > 0 {
			return fmt.Sprintf("https://www.twitter.com/%s/", a.Twitter)
		}
	}

	return ""
}

func (a *Associate) getLine(section string, bold bool) string {
	line := fmt.Sprintf("%s%s%s",
		a.FirstName,
		func() string {
			if a.PreferredName == "" {
				return " "
			}
			return fmt.Sprintf(" (%s) ", a.PreferredName)
		}(),
		a.LastName)

	// Pick the first available social media link from the defined order.
	for i := 0; i < 6; i++ {
		if str := a.getLink(i); len(str) > 0 {
			line = fmt.Sprintf("[%s](%s)", line, str)
			break
		}
	}

	// Reformat the line and write it.
	line = fmt.Sprintf("%s, %s", line, a.Position)
	if bold && section != alm && a.isExec() {
		line = "**" + line + "**"
	}
	line = "- " + line

	return line
}

func (a *Associate) isExec() bool {
	return strings.Index(a.Position, "VP") >= 0 || strings.Index(a.Position, "President") >= 0
}

func (a *Associate) hasGraduated() bool {
	return a.Graduated >= 0
}
