package project

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const activeProjectTableBasePath = "./assets/projects_active.md"
const pastProjectTableBasePath = "./assets/projects_past.md"

// MakeTable creates a list of project lines.
func MakeTable(projects []Project, active bool) []string {

	if len(projects) == 0 {
		return []string{}
	}

	lines := helpers.ReadFileBase(func() string {
		if active {
			return activeProjectTableBasePath
		}
		return pastProjectTableBasePath
	}(), len(projects), 4)

	for _, proj := range projects {
		projListing := fmt.Sprintf("|[%s](%s)|%s|%s|%s|[%s](/team/%s)|%s|%s|",
			proj.Title,
			proj.Link,
			helpers.TablePadder,
			proj.Description,
			helpers.TablePadder,
			proj.Department,
			helpers.StringToFileName(proj.Department),
			helpers.TablePadder,
			proj.Instructions,
		)
		lines = append(lines, projListing)
	}

	return lines

}
