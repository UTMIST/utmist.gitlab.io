package project

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const activeProjectHeader = "## **Active Projects**"
const pastProjectHeader = "## **Past Projects**"
const projectFilename = "content/projects.md"
const projectPhotoLink = "![Projects](/images/projects.png)"

// MakeList creates a list of project lines.
func MakeList(projects *[]Project, active, deptPage bool) []string {

	if len(*projects) == 0 {
		return []string{}
	}

	lines := []string{func() string {
		if active {
			return activeProjectHeader
		}
		return pastProjectHeader
	}()}

	for _, proj := range *projects {
		lines = append(lines, fmt.Sprintf("##### **%s**",
			func() string {
				if len(proj.Link) == 0 {
					return proj.Title
				}
				return fmt.Sprintf("[%s](/projects/%s)", proj.Title, proj.Link)
			}()))

		if len(proj.Description) > 0 {
			lines = append(lines, proj.Description)
		}
		if len(proj.Department) > 0 && !deptPage {
			lines = append(lines,
				fmt.Sprintf("- _Department_: [%s](/team/%s)",
					proj.Department,
					helpers.StringToFileName(proj.Department)))
		}
		if len(proj.Instructions) > 0 && proj.Status == ActiveStatus {
			lines = append(lines,
				fmt.Sprintf("- _Joining_: %s", proj.Instructions))
		}

	}

	return lines

}

// GeneratePages generates a page for the project list.
func GeneratePages(projects, pastProjects *[]Project, desc string) {
	lines := append(helpers.GenerateHeader("Projects", "0001-01-03"),
		projectPhotoLink, "", helpers.Breakline)
	if len(desc) > 0 {
		lines = append(lines, "", helpers.Breakline, desc)
	}

	// Load lists of active/past projects.
	lines = append(lines, MakeList(projects, true, false)...)
	if len(*projects) > 0 && len(*pastProjects) > 0 {
		lines = append(lines, helpers.Breakline)
	}
	lines = append(lines, MakeList(pastProjects, false, false)...)

	helpers.OverwriteWithLines(projectFilename, lines)
}
