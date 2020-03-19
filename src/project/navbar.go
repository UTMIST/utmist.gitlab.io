package project

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// Where the department list starts in the config.
const start = "    - title: Projects"

// GenerateNavbarProjectLinks generates project links for the navbar dropdown.
func GenerateNavbarProjectLinks(projects *[]Project, lines *[]string) {
	projectTitles := []string{}
	projectURLs := []string{}
	for _, proj := range *projects {
		projectTitles = append(projectTitles, proj.Title)
		projectURLs = append(projectURLs, proj.Link)
	}

	helpers.StitchExternalLink(lines, projectTitles, projectURLs, start)
}
