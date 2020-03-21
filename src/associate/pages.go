package associate

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// Generate a page for the a department.
func generateDepartmentPage(
	title string,
	associates []Associate,
	positions []position.Position,
	projects []project.Project,
	pastProjects []project.Project) {

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
	assocLines, execLines := []string{}, []string{}
	for _, associate := range associates {
		if associate.isExec() && !associate.hasGraduated() {
			execLines = append(execLines, associate.getLine(title, true, true))
		} else {
			assocLines = append(assocLines, associate.getLine(title, true, true))
		}
	}

	// Stitch the new lines back in with any existing open positions.
	lines = append(lines, execLines...)
	lines = append(lines, assocLines...)
	if len(projects)+len(pastProjects) > 0 {
		lines = append(lines, helpers.Breakline)
	}
	lines = append(lines, project.MakeList(&projects, true)...)
	lines = append(lines, project.MakeList(&pastProjects, false)...)
	lines = append(lines, position.MakeList(&positions, true)...)

	// Write to the new file path.
	filepath := fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(title))
	helpers.OverwriteWithLines(filepath, lines)
}

// GenerateAssociatePages generates all the department pages.
func GenerateAssociatePages(
	associates *[]Associate,
	positions *[]position.Position,
	projects *[]project.Project,
	pastProjects *[]project.Project) {
	logger.GenerateGroupLog("associate")

	// Populate the departments with empty associate/position/project lists.
	deptAssocMap := GroupByDept(associates)
	deptPosMap := position.GroupByDept(positions)
	deptProjMap := project.GroupByDept(projects)
	deptPastProjMap := project.GroupByDept(pastProjects)

	// Generate each department page.
	os.Mkdir(assocDirPath, os.ModePerm)

	for deptName, deptAssociates := range deptAssocMap {
		sort.Sort(List(deptAssociates))

		generateDepartmentPage(
			deptName,
			deptAssociates,
			deptPosMap[deptName],
			deptProjMap[deptName],
			deptPastProjMap[deptName])
	}
}
