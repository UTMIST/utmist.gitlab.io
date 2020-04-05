package department

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// Paths for the event files.
const assocDirPath = "./content/team/"

// GeneratePage generates a page for the a department.
func GeneratePage(
	title string,
	associates []associate.Associate,
	description string,
	positions []position.Position,
	projects []project.Project,
	pastProjects []project.Project) {

	helpers.GenerateLog(title)

	// Get page title and date and generate the header.
	displayTitle, yearStr := func() (string, string) {
		if title == helpers.ALM {
			return fmt.Sprintf("Our %s", helpers.ALM), "0000-01-02"
		}
		return title, "0000-01-01"
	}()
	lines := append(
		helpers.GenerateFrontMatter(
			displayTitle, yearStr, "", []string{"Team"}),
		[]string{description, ""}...)

	// Write a list entry for every member.
	assocLines, execLines := []string{}, []string{}
	for _, associate := range associates {
		if associate.IsExec() && !associate.HasGraduated() {
			execLines = append(execLines,
				associate.GetLine(title, true, true))
		} else {
			assocLines = append(assocLines,
				associate.GetLine(title, true, true))
		}
	}

	// Stitch the new lines back in with projects and positions.
	lines = append(lines, execLines...)
	lines = append(lines, assocLines...)
	if len(projects)+len(pastProjects) > 0 {
		lines = append(lines, helpers.Breakline)
	}
	lines = append(lines, project.MakeList(&projects, true, true)...)
	lines = append(lines, project.MakeList(&pastProjects, false, true)...)
	lines = append(lines, position.MakeList(&positions, true, "")...)

	// Write to the new file path.
	filepath := fmt.Sprintf("%s%s.md", assocDirPath, strings.ToLower(title))
	helpers.OverwriteWithLines(filepath, lines)
}

// GeneratePages generates all the department pages.
func GeneratePages(
	associates *[]associate.Associate,
	descriptions *map[string]string,
	positions *[]position.Position,
	projects *[]project.Project,
	pastProjects *[]project.Project) {

	helpers.GenerateGroupLog("associate")

	// Populate the departments with associate/position/project maps.
	deptAssocMap := associate.GroupByDept(associates)
	deptPosMap := position.GroupByDept(positions)
	deptProjMap := project.GroupByDept(projects)
	deptPastProjMap := project.GroupByDept(pastProjects)

	os.Mkdir(assocDirPath, os.ModePerm)

	// Generate each department page.
	for deptName, deptAssociates := range deptAssocMap {
		sort.Sort(associate.List(deptAssociates))

		description, exists := (*descriptions)[deptName]
		if !exists {
			description = ""
		}

		GeneratePage(
			deptName,
			deptAssociates,
			description,
			deptPosMap[deptName],
			deptProjMap[deptName],
			deptPastProjMap[deptName])
	}
}
