package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const associateListSubstitution = "[//]: # associates"
const departmentListSubstitution = "[//]: # departments"
const executiveListSubstitution = "[//]: # executives"

// GenerateDepartmentAssociateLists inserts generated lists of associates into dept pages.
func GenerateDepartmentAssociateLists(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry) {

	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		depts := strings.Split(os.Getenv(fmt.Sprintf("DEPTS_%d", y)), ",")
		deptToEntryMap := map[string][]associate.Entry{}
		for _, dept := range depts {
			deptToEntryMap[dept] = []associate.Entry{}
		}
		for _, entry := range (*entries)[y] {
			if _, exists := deptToEntryMap[entry.Department]; !exists {
				continue
			}
			deptToEntryMap[entry.Department] =
				append(deptToEntryMap[entry.Department], entry)
		}

		for dept, deptEntries := range deptToEntryMap {
			filepath := helpers.RelativeFilePath(y, lastYear, dept)
			if _, err := os.Stat(filepath); err != nil {
				log.Println(err)
				continue
			}

			lines := helpers.ReadContentLines(filepath)
			associateLines := associate.MakeEntryList(associates, &deptEntries, true)
			lines = helpers.SubstituteString(
				lines,
				associateLines,
				associateListSubstitution)

			yearLine := getYearListString(dept, firstYear, lastYear, y)
			lines = helpers.SubstituteString(
				lines,
				[]string{yearLine},
				yearListSubstitution)
			helpers.OverwriteWithLines(filepath, lines)
		}
	}
}

// GenerateTeamDepartmentList inserts generated lists of executives into team pages.
func GenerateTeamDepartmentList(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry) {

	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {

		depts := strings.Split(os.Getenv(fmt.Sprintf("DEPTS_%d", y)), ",")
		deptToEntryMap := map[string][]associate.Entry{}
		for _, dept := range depts {
			deptToEntryMap[dept] = []associate.Entry{}
		}

		filepath := helpers.RelativeFilePath(y, lastYear, "team")
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		lines := helpers.ReadContentLines(filepath)
		deptLines := associate.MakeDepartmentList(&depts)
		lines = helpers.SubstituteString(
			lines,
			deptLines,
			departmentListSubstitution)

		yearLine := getYearListString("team", firstYear, lastYear, y)
		lines = helpers.SubstituteString(
			lines,
			[]string{yearLine},
			yearListSubstitution)
		helpers.OverwriteWithLines(filepath, lines)
	}
}

// GenerateTeamExecutiveList inserts generated lists of executives into team pages.
func GenerateTeamExecutiveList(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry) {

	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		execs := []associate.Entry{}
		for _, entry := range (*entries)[y] {
			if entry.IsExecutive() {
				execs = append(execs, entry)
			}
		}

		filepath := helpers.RelativeFilePath(y, lastYear, "team")
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		lines := helpers.ReadContentLines(filepath)
		newLines := associate.MakeEntryList(associates, &execs, false)
		lines = helpers.SubstituteString(
			lines,
			newLines,
			executiveListSubstitution)
		helpers.OverwriteWithLines(filepath, lines)
	}
}
