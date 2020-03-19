package associate

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// Generate a page for the a department.
func generateDepartmentPage(
	title string,
	associates *[]Associate,
	positions *[]position.Position) {

	logger.GenerateLog(title)

	// Get page title and date and generate the header.
	displayTitle, yearStr := func() (string, string) {
		if title == alm {
			return fmt.Sprintf("Our %s", alm), "0000-01-02"
		}
		return title, "0000-01-01"
	}()
	lines := helpers.GenerateHugoPageHeader(
		displayTitle, yearStr, "", []string{"Team"})

	// Write a list entry for every member.
	assocLines := []string{}
	execLines := []string{}
	for _, associate := range *associates {
		if associate.isExec() && !associate.hasGraduated() {
			execLines = append(execLines, associate.getLine(title, true))
		} else {
			assocLines = append(assocLines, associate.getLine(title, true))
		}
	}

	// Stitch the new lines back in with any existing open positions.
	lines = append(lines, execLines...)
	lines = append(lines, assocLines...)
	lines = append(lines, position.MakeList(positions, true)...)

	// Write to the new file path.
	filepath := fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(title))
	helpers.OverwriteWithLines(filepath, lines)
}

// GenerateAssociatePages generates all the department pages.
func GenerateAssociatePages(associates *[]Associate, positions *[]position.Position) {
	logger.GenerateGroupLog("associate")

	// Populate the departments with empty associate list.
	departments := map[string][]Associate{}
	deptPositions := map[string][]position.Position{}
	for _, dept := range GetDepartmentNames() {
		departments[dept] = []Associate{}
		deptPositions[dept] = []position.Position{}
	}

	// Load associates into their department's associate list.
	for _, associate := range *associates {
		if deptList, exists := departments[associate.Department]; exists {
			departments[associate.Department] = append(deptList, associate)
		}
	}

	// Load positions into their department's associate list.
	for _, position := range *positions {
		if posList, exists := deptPositions[position.Department]; exists {
			deptPositions[position.Department] = append(posList, position)
		}
	}

	// Generate each department page.
	os.Mkdir(assocDirPath, os.ModePerm)
	for deptName, deptAssociates := range departments {
		sort.Sort(List(deptAssociates))

		generateDepartmentPage(deptName, &deptAssociates, func() *[]position.Position {
			if pos, exists := deptPositions[deptName]; exists {
				return &pos
			}
			return nil
		}())
	}
}
