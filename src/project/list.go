package project

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const activeProjectListBasePath = "./assets/projects_active.md"
const pastProjectListBasePath = "./assets/projects_past.md"

const projectCopyFilename = "assets/projects.md"
const projectFilename = "content/projects.md"

// MakeList creates a list of project lines.
func MakeList(projects *[]Project, active, deptPage bool) []string {

	if len(*projects) == 0 {
		return []string{}
	}

	lines := helpers.ReadFileBase(func() string {
		if active {
			return activeProjectListBasePath
		}
		return pastProjectListBasePath
	}(), len(*projects), 1)

	for _, proj := range *projects {
		head := fmt.Sprintf("##### **%s**",
			func() string {
				if len(proj.Link) == 0 {
					return proj.Title
				}
				return fmt.Sprintf("[%s](%s)", proj.Title, proj.Link)
			}())

		dept := func() string {
			if deptPage {
				return ""
			}
			return fmt.Sprintf("- _Department_: [%s](%s)",
				proj.Department,
				helpers.StringToFileName(proj.Department))
		}()
		desc := fmt.Sprintf("- _Description_: %s", proj.Description)
		join := func() string {
			if proj.Status != ActiveStatus {
				return ""
			}
			return fmt.Sprintf("- _Joining_: %s", proj.Instructions)
		}()

		lines = append(lines, []string{head, dept, desc, join}...)

	}

	return lines

}

// GenerateProjectListPage generates a page for the project list.
func GenerateProjectListPage(projects, pastProjects *[]Project) {
	lines := helpers.ReadContentLines(projectCopyFilename)

	// Load lists of active/past projects.
	lines = append(lines, MakeList(projects, true, false)...)
	if len(*projects) > 0 && len(*pastProjects) > 0 {
		lines = append(lines, helpers.Breakline)
	}
	lines = append(lines, MakeList(pastProjects, false, false)...)

	helpers.OverwriteWithLines(projectFilename, lines)
}
