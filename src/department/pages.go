package department

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// GenerateDeptPage generates a page for the a department.
func GenerateDeptPage(
	title string,
	associates []associate.Associate,
	description string,
	positions []position.Position,
	projects []project.Project,
	pastProjects []project.Project) {

	logger.GenerateLog(title)

	// Get page title and date and generate the header.
	displayTitle, yearStr := func() (string, string) {
		if title == associate.ALM {
			return fmt.Sprintf("Our %s", associate.ALM), "0000-01-02"
		}
		return title, "0000-01-01"
	}()
	lines := append(helpers.GenerateHugoPageHeader(
		displayTitle, yearStr, "", []string{"Team"}), description)

	// Write a list entry for every member.
	assocLines, execLines := []string{}, []string{}
	for _, associate := range associates {
		if associate.IsExec() && !associate.HasGraduated() {
			execLines = append(execLines, associate.GetLine(title, true, true))
		} else {
			assocLines = append(assocLines, associate.GetLine(title, true, true))
		}
	}

	// Stitch the new lines back in with any existing open positions.
	lines = append(lines, execLines...)
	lines = append(lines, assocLines...)
	if len(projects)+len(pastProjects) > 0 {
		lines = append(lines, helpers.Breakline)
	}
	lines = append(lines, project.MakeList(&projects, true, true)...)
	lines = append(lines, project.MakeList(&pastProjects, false, true)...)
	lines = append(lines, position.MakeList(&positions, true)...)

	// Write to the new file path.
	filepath := fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(title))
	helpers.OverwriteWithLines(filepath, lines)
}

// GenerateDeptPages generates all the department pages.
func GenerateDeptPages(
	associates *[]associate.Associate,
	descriptions *map[string]string,
	positions *[]position.Position,
	projects *[]project.Project,
	pastProjects *[]project.Project) {

	logger.GenerateGroupLog("associate")

	// Populate the departments with empty associate/position/project lists.
	deptAssocMap := associate.GroupByDept(associates)
	deptPosMap := position.GroupByDept(positions)
	deptProjMap := project.GroupByDept(projects)
	deptPastProjMap := project.GroupByDept(pastProjects)

	// Generate each department page.
	os.Mkdir(assocDirPath, os.ModePerm)

	for deptName, deptAssociates := range deptAssocMap {
		sort.Sort(associate.List(deptAssociates))

		description, exists := (*descriptions)[deptName]
		if !exists {
			description = ""
		}

		GenerateDeptPage(
			deptName,
			deptAssociates,
			description,
			deptPosMap[deptName],
			deptProjMap[deptName],
			deptPastProjMap[deptName])
	}
}
