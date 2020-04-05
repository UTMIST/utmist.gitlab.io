package associate

import (
	"fmt"
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
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

	// Compare positions, then last names, then first names.
	for _, criteria := range []int{
		strings.Compare(a[i].Position, a[j].Position),
		strings.Compare(a[i].LastName, a[j].LastName),
		strings.Compare(a[i].FirstName, a[j].FirstName)} {
		switch criteria {
		case -1:
			return true
		case 1:
			return false
		}
	}
	return false
}

// Method Swap() to implement sort.Sort.
func (a List) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

const website = ""
const linkedin = "https://linkedin.com/in/"
const github = "https://www.github.com/"
const gitlab = "https://www.gitlab.com/"
const facebook = "https://www.facebook.com/"
const twitter = "https://www.twitter.com/"

// Return personal link for associate.
func (a *Associate) getLink() string {

	// Order of links.
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

// GetLine creates a line entry for associate.
func (a *Associate) GetLine(section string, bold, list bool) string {

	// Set up associate's name.
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
	if section == helpers.ALM {
		line = fmt.Sprintf("%s, %s", line, a.Discipline)
	} else {
		line = fmt.Sprintf("%s, %s", line, a.Position)
		if bold && a.IsExec() {
			line = "**" + line + "**"
		}
	}

	if !list {
		return line
	}
	return "- " + line
}

// IsExec returns whether this associate is an executive member.
func (a *Associate) IsExec() bool {
	return strings.Index(a.Position, "VP") >= 0 ||
		strings.Index(a.Position, "President") >= 0
}

// HasGraduated returns whether this associate has graduated or left.
func (a *Associate) HasGraduated() bool {
	return 0 <= a.Graduated && a.Graduated < time.Now().Year()
}

// GroupByDept groups associates into their own department list.
func GroupByDept(associates *[]Associate) map[string][]Associate {

	// Populate an empty list for every department.
	deptAssociates := map[string][]Associate{}
	for _, dept := range helpers.GetDeptNames(false) {
		deptAssociates[dept] = []Associate{}
	}

	// Insert associates into their appropriate department, if it exists.
	for _, assoc := range *associates {
		assocList, exists := deptAssociates[assoc.Department]
		if !exists {
			continue
		}
		deptAssociates[assoc.Department] = append(assocList, assoc)
	}
	return deptAssociates
}

// Dept implements hasDepartment().
func (a Associate) Dept() string {
	return a.Department
}
