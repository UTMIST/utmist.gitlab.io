package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const departmentListSubstitution = "[//]: # departments"
const executiveListSubstitution = "[//]: # executives"

// GenerateTeamDepartmentList inserts generated lists of departments into team pages.
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
		deptLines := associate.MakeDepartmentList(&depts, lastYear, y)
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
