package associate

import (
	"fmt"
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// Associate represents an associateutive member's database row.
type Associate struct {
	FirstName     string
	PreferredName string
	LastName      string

	UofTEmail   string
	Email       string
	PhoneNumber string

	Position   string
	Department string
	Retirement time.Time
	Discipline string

	ProfilePicture string
	Website        string
	LinkedIn       string
	GitLab         string
	GitHub         string
	Facebook       string
	Twitter        string
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
	bases := []string{
		website, linkedin, gitlab, github, facebook, twitter}
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

// GetEntry creates a line entry for associate.
func (a *Associate) GetEntry(section string, bold, list bool) string {

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

	if section == helpers.ALM {
		return fmt.Sprintf("- **%s, %s**, %s", line, a.Discipline, a.Position)
	}

	line = fmt.Sprintf("%s, %s", line, strings.Split(a.Position, " (")[0])
	if a.IsExec() && bold || section == helpers.ALM {
		line = fmt.Sprintf("**%s**", line)
	}
	line = fmt.Sprintf("- %s", line)

	return line
}

// IsExec returns whether this associate is an executive member.
func (a *Associate) IsExec() bool {
	return strings.Index(a.Position, "VP") >= 0 ||
		strings.Index(a.Position, "President") >= 0
}

// HasRetired returns whether this associate has graduated or left.
func (a *Associate) HasRetired() bool {
	return a.Retirement.Before(time.Now())
}
