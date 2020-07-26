package project

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const activeProjectHeader = "## **Active Projects**"
const pastProjectHeader = "### Past Projects"
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
			deptLine := "- _Department(s)_: "
			for _, dept := range proj.Department {
				deptLine = fmt.Sprintf("%s [%s](/%s),",
					deptLine, dept,
					helpers.StringToSimplePath(dept))
			}
			lines = append(lines,
				fmt.Sprintf("%s.", deptLine[:len(deptLine)-1]))
		}

		if len(proj.Instructions) > 0 && proj.Status == ActiveStatus {
			lines = append(lines,
				fmt.Sprintf("- _Joining_: %s", proj.Instructions))
		}
	}

	return append([]string{"", helpers.Breakline}, lines...)
}

// GenerateList generates the project list.
func GenerateList(projects, pastProjects *[]Project, desc string) {
	lines := append(helpers.GenerateHeader("Projects", "0001-01-03"), desc)

	// Load lists of active/past projects.
	lines = append(lines, MakeList(projects, true, false)...)
	lines = append(lines, MakeList(pastProjects, false, false)...)

	helpers.OverwriteWithLines(projectFilename, lines)
}
