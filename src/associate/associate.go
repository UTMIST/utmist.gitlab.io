package associate

import (
	"fmt"
	"strings"
)

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

	// First compare positions.
	switch strings.Compare(a[i].Position, a[j].Position) {
	case -1:
		return true
	case 1:
		return false
	}

	// Then compare last names.
	switch strings.Compare(a[i].LastName, a[j].LastName) {
	case -1:
		return true
	case 1:
		return false
	}

	// Finally compare first names.
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

const website = ""
const linkedin = "https://linkedin.com/in/"
const gitlab = "https://www.github.com/"
const github = "https://www.gitlab.com/"
const facebook = "https://www.facebook.com/"
const twitter = "https://www.twitter.com/"

// Return personal link for associate.
func (a *Associate) getLink() string {

	bases := []string{website, linkedin, gitlab, github, facebook, twitter}
	links := []string{
		a.Website, a.LinkedIn, a.GitLab, a.GitHub, a.Facebook, a.Twitter}

	// Return the first link found.
	for i, link := range links {
		if len(link) == 0 {
			continue
		}
		return fmt.Sprintf("%s%s/", bases[i], links[i])
	}

	return ""
}

// Create line entry for associate.
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
	if str := a.getLink(); len(str) > 0 {
		line = fmt.Sprintf("[%s](%s)", line, str)
	}

	// Reformat the line and write it. List just graduation on alumni page.
	if section == alm {
		line = fmt.Sprintf("%s, %s", line, a.Discipline)
	} else {
		line = fmt.Sprintf("%s, %s", line, a.Position)
		if bold && a.isExec() {
			line = "**" + line + "**"
		}
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
