package associate

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"

	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// Generate a page for the a department.
func generateDepartmentPage(name string, associates []Associate) {
	logger.GenerateLog(fmt.Sprintf("%s department", name))

	lines := hugo.GeneratePageHeader(fmt.Sprintf("%s Department", name), "0001-01-01", "", []string{"Team"})

	// Write a list entry for every member; skip the alumni (retired).
	for _, associate := range associates {
		if associate.Retired >= 0 {
			continue
		}

		line := fmt.Sprintf("%s%s%s",
			associate.FirstName,
			func() string {
				if associate.PreferredName == "" {
					return " "
				}
				return fmt.Sprintf(" (%s) ", associate.PreferredName)
			}(),
			associate.LastName)

		// Pick the first available social media link from the defined order.
		for i := 0; i < 6; i++ {
			if str := associate.getLink(i); len(str) > 0 {
				line = fmt.Sprintf("[%s](%s)", line, str)
				break
			}
		}

		// Reformat the line and write it.
		line = fmt.Sprintf("%s, %s", line, associate.Position)
		if strings.Index(associate.Position, "VP") >= 0 ||
			strings.Index(associate.Position, "President") >= 0 {
			line = "**" + line + "**"
		}
		line = "- " + line
		lines = append(lines, line)
	}

	filename := fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(name))
	helpers.OverwriteWithLines(filename, lines)
}

// GenerateAssociatePages generates all the department pages.
func GenerateAssociatePages(associates []Associate) {
	logger.GenerateGroupLog("associate")

	// Populate the departments with empty associate list.
	depts := map[string][]Associate{}
	for _, dept := range getDepartments() {
		depts[dept] = []Associate{}
	}

	// Load associates into their department's associate list.
	for _, associate := range associates {
		for _, dept := range associate.Departments {
			if deptList, exists := depts[dept]; exists {
				depts[dept] = append(deptList, associate)
			}
		}
	}

	// Generate each department page.
	os.Mkdir(assocDirPath, os.ModePerm)
	for deptName, deptAssociates := range depts {
		generateDepartmentPage(deptName, deptAssociates)
	}
}
