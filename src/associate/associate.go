package associate

import (
	"fmt"
)

const website = ""
const linkedin = "https://linkedin.com/in/"
const github = "https://www.github.com/"
const gitlab = "https://www.gitlab.com/"
const facebook = "https://www.facebook.com/"
const twitter = "https://www.twitter.com/"

// Associate represents an associateutive member's database row.
type Associate struct {
	GivenName     string
	PreferredName string
	Surname       string

	UofTEmail   string
	OtherEmail  string
	PhoneNumber string
	Discipline  string

	ProfilePicture string
	Website        string
	LinkedIn       string
	GitHub         string
	GitLab         string
	Facebook       string
	Twitter        string
}

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

// Return formatted full name for associate.
func (a *Associate) getName() string {
	if len(a.PreferredName) > 0 {
		return fmt.Sprintf("%s (%s) %s", a.GivenName, a.PreferredName, a.Surname)
	}
	return fmt.Sprintf("%s %s", a.GivenName, a.Surname)
}

// Return target link for associate
func (a *Associate) getTargetLink(linkType string) string {
	link := ""
	switch linkType {
	case "linkedin":
		if a.LinkedIn != "" {
			link = fmt.Sprintf("%s%s/", linkedin, a.LinkedIn)
		}
	case "github":
		if a.GitHub != "" {
			link = fmt.Sprintf("%s%s/", github, a.GitHub)
		}
	case "gitlab":
		if a.GitLab != "" {
			link = fmt.Sprintf("%s%s/", gitlab, a.GitLab)
		}
	case "twitter":
		if a.Twitter != "" {
			link = fmt.Sprintf("%s%s/", twitter, a.Twitter)
		}
	case "facebook":
		if a.Facebook != "" {
			link = fmt.Sprintf("%s%s/", facebook, a.Facebook)
		}
	case "personal":
		if a.Website != "" {
			link = fmt.Sprintf("%s%s/", website, a.Website)
		}
	}

	return link
}
