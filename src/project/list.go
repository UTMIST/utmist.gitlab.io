package project

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const activeProjectTableBasePath = "./assets/projects_active.md"
const pastProjectTableBasePath = "./assets/projects_past.md"

const projectCopyFilename = "assets/projects.md"
const projectFilename = "content/projects.md"

// MakeList creates a list of project lines.
func MakeList(projects *[]Project, active bool) []string {

	if len(*projects) == 0 {
		return []string{}
	}

	lines := helpers.ReadFileBase(func() string {
		if active {
			return activeProjectTableBasePath
		}
		return pastProjectTableBasePath
	}(), len(*projects), 4)

	for _, proj := range *projects {
		projListing := fmt.Sprintf("|[%s](%s)|%s|%s|%s|%s|",
			proj.Title,
			proj.Link,
			helpers.TablePadder,
			proj.Description,
			helpers.TablePadder,
			proj.Instructions,
		)
		lines = append(lines, projListing)
	}

	return lines

}

// GenerateProjectListPage generates a page for the project list.
func GenerateProjectListPage(projects, pastProjects *[]Project) {
	lines := helpers.ReadContentLines(projectCopyFilename)

	// Load tables of active/past projects.
	lines = append(lines, MakeList(projects, true)...)
	lines = append(lines, MakeList(pastProjects, false)...)

	helpers.OverwriteWithLines(projectFilename, lines)
}
