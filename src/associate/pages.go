package associate

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// Generate a page for the a department.
func generateDepartmentPage(title string, associates []Associate) {
	logger.GenerateLog(fmt.Sprintf("%s department", title))

	// Get page title and date and generate the header.
	title, yearStr := func() (string, string) {
		if title == alm {
			return fmt.Sprintf("Our %s", alm), "0000-01-02"
		}
		return title, "0000-01-01"
	}()
	lines := helpers.GenerateHugoPageHeader(
		title, yearStr, "", []string{"Team"})

	// Write a list entry for every member.
	assocLines := []string{}
	execLines := []string{}
	for _, associate := range associates {
		if associate.isExec() && !associate.hasGraduated() {
			execLines = append(execLines, associate.getLine(title, true))
		} else {
			assocLines = append(assocLines, associate.getLine(title, true))
		}
	}

	// Stitch the new lines back in.
	lines = append(lines, execLines...)
	lines = append(lines, assocLines...)

	// Write to the new file path.
	filepath := fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(title))
	helpers.OverwriteWithLines(filepath, lines)
}

// GenerateAssociatePages generates all the department pages.
func GenerateAssociatePages(associates []Associate) {
	logger.GenerateGroupLog("associate")

	// Populate the departments with empty associate list.
	depts := map[string][]Associate{}
	for _, dept := range GetDepartmentNames() {
		depts[dept] = []Associate{}
	}

	// Load associates into their department's associate list.
	for _, associate := range associates {
		if deptList, exists := depts[associate.Department]; exists {
			depts[associate.Department] = append(deptList, associate)
		}
	}

	// Generate each department page.
	os.Mkdir(assocDirPath, os.ModePerm)
	for deptName, deptAssociates := range depts {
		sort.Sort(List(deptAssociates))
		generateDepartmentPage(deptName, deptAssociates)
	}
}
