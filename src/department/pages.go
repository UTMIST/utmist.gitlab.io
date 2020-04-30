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
const alumniDeptTitle = "Our Alumni & Past Members"

// GeneratePage generates a page for the a department.
func GeneratePage(
	title string,
	associates []associate.Associate,
	descriptions *map[string]string,
	positions []position.Position,
	projects []project.Project,
	pastProjects []project.Project) {

	helpers.GenerateLog(title)

	// Get page title and date and generate the header.
	displayTitle, yearStr := title, "0000-01-01"
	if title == helpers.ALM {
		displayTitle, yearStr = alumniDeptTitle, "0000-01-02"
	}

	lines := append(
		helpers.GenerateHeader(displayTitle, yearStr),
		[]string{(*descriptions)[title], ""}...)

	regularPage := title != helpers.ADV && title != helpers.ADM
	fmt.Println(regularPage)

	// Write a list entry for every member.
	assocLines, execLines := []string{}, []string{}
	for _, associate := range associates {
		if associate.IsExec(true) && !associate.HasRetired() {
			execLines = append(execLines,
				associate.GetEntry(title, regularPage, true))
		} else {
			assocLines = append(assocLines,
				associate.GetEntry(title, regularPage, true))
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
	lines = append(lines, position.MakeList(
		&positions, true, "", (*descriptions)["Recruitment"])...)

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

		GeneratePage(
			deptName,
			deptAssociates,
			descriptions,
			deptPosMap[deptName],
			deptProjMap[deptName],
			deptPastProjMap[deptName])
	}
}
