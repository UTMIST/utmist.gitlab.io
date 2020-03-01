package associate

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// Generate a page for the a department.
func generateDepartmentPage(name string, associates []Associate) {
	logger.GenerateLog(fmt.Sprintf("%s team", name))

	// Create file for the page and write the header.
	f, err := os.Create(fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(name)))
	if err != nil {
		logger.GenerateErrorLog(fmt.Sprintf("%s team", name))
	}
	defer f.Close()
	hugo.GeneratePageHeader(f, fmt.Sprintf("%s Department", name), "0001-01-01", "", []string{"Team"})

	// Write a list entry for every member; skip the alumni (retired).
	for _, associate := range associates {
		if associate.Retired >= 0 {
			continue
		}

		var line string
		if associate.PreferredName != "" {
			line = fmt.Sprintf("%s (%s) %s",
				associate.FirstName,
				associate.PreferredName,
				associate.LastName)
		} else {
			line = fmt.Sprintf("%s %s",
				associate.FirstName,
				associate.LastName)
		}

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
		fmt.Fprintln(f, line)
	}

	// Try closing the file.
	if err := f.Close(); err != nil {
		logger.GenerateErrorLog(fmt.Sprintf("%s team", name))
	}

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
