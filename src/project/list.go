package project

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const activeProjectHeader = "## **Active Projects**"
const pastProjectHeader = "## **Past Projects**"
const projectFilename = "content/projects.md"

// MakeList creates a list of project lines.
func MakeList(projects *[]Project, active, deptPage bool) []string {

	var lines []string
	if len(*projects) == 0 {
		return lines
	} else if active {
		lines = append(lines, activeProjectHeader)
	} else {
		lines = append(lines, pastProjectHeader)
	}

	for _, proj := range *projects {
		title := proj.Title
		if len(proj.Link) > 0 {
			title = fmt.Sprintf("[%s](%s)", proj.Title, proj.Link)
		}

		lines = append(lines, fmt.Sprintf("##### **%s**", title))
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

// GenerateList generates the project list.
func GenerateList(projects, pastProjects *[]Project, desc string) {
	lines := append(helpers.GenerateHeader("Projects", "0001-01-03"),
		desc, "", helpers.Breakline)

	// Load lists of active/past projects.
	lines = append(lines, MakeList(projects, true, false)...)
	if len(*projects) > 0 && len(*pastProjects) > 0 {
		lines = append(lines, helpers.Breakline)
	}
	lines = append(lines, MakeList(pastProjects, false, false)...)

	helpers.OverwriteWithLines(projectFilename, lines)
}
