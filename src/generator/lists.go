package generator

import (
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// GenerateDepartmentAssociateLists inserts generated lists of associates into dept pages.
func GenerateDepartmentAssociateLists(
	associates *map[string]associate.Associate,
	entries *map[int]map[string][]associate.Entry,
	dept string,
	year int) []string {

	yearDeptEntries := (*entries)[year][dept]
	return associate.MakeEntryList(associates, &yearDeptEntries, true)
}

// GenerateEventList inserts generated lists of events into the events page.
func GenerateEventList(
	eventMap *map[int][]event.Event,
	year int) []string {

	events := (*eventMap)[year]
	return event.GenerateListPage(&events)
}

// GeneratePositionList generates position lists for the recruitment page.
func GeneratePositionList(
	positions *map[string][]position.Position,
	level string) []string {

	levelPositions := (*positions)[level]
	return position.MakeList(
		&levelPositions,
		false,
		level)
}

// GenerateDeptPositionLists generates position lists for department pages.
func GenerateDeptPositionLists(
	positions *map[string][]position.Position,
	department string) []string {

	departmentPositions := (*positions)[department]
	return position.MakeList(
		&departmentPositions,
		false,
		"Associate/Executive")
}

// GenerateProjectLists generates projects lists for the projects page.
func GenerateProjectLists(
	projectMap *map[int]map[string][]project.Project,
	projectType string,
	year int) []string {

	projects := (*projectMap)[year][projectType]
	return project.MakeList(&projects)
}

// GenerateTeamDepartmentList inserts generated lists of departments into team pages.
func GenerateTeamDepartmentList(
	entries *map[int]map[string][]associate.Entry,
	year int) []string {

	departments := []string{}
	for dept := range (*entries)[year] {
		departments = append(departments, dept)
	}

	_, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	return associate.MakeDepartmentList(&departments, lastYear, year)

}

// GenerateTeamExecutiveList inserts generated lists of executives into team pages.
func GenerateTeamExecutiveList(
	associates *map[string]associate.Associate,
	entries *map[int]map[string][]associate.Entry,
	year int) []string {

	execs := []associate.Entry{}
	for _, deptEntries := range (*entries)[year] {
		for _, entry := range deptEntries {
			if entry.IsExecutive() {
				execs = append(execs, entry)
			}
		}
	}

	return associate.MakeEntryList(associates, &execs, false)
}
