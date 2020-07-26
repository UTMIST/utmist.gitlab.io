package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const associatesSubstitution = "[//]: # associates"
const executivesSubstitution = "[//]: # executives"

// GenerateDepartmentAssociateLists inserts generated lists of associates into dept pages.
func GenerateDepartmentAssociateLists(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry) {

	firstYear, lastYear := helpers.GetYearRange()
	for year := firstYear; year <= lastYear; year++ {
		depts := strings.Split(os.Getenv(fmt.Sprintf("DEPTS_%d", year)), ",")
		deptToEntryMap := map[string][]associate.Entry{}
		for _, dept := range depts {
			deptToEntryMap[dept] = []associate.Entry{}
		}
		for _, entry := range (*entries)[year] {
			if _, exists := deptToEntryMap[entry.Department]; !exists {
				continue
			}
			deptToEntryMap[entry.Department] =
				append(deptToEntryMap[entry.Department], entry)
		}

		for dept, deptEntries := range deptToEntryMap {
			filepath := helpers.RelativeFilePath(year, lastYear, dept)
			if _, err := os.Stat(filepath); err != nil {
				log.Println(err)
				continue
			}

			lines := helpers.ReadContentLines(filepath)
			newLines := associate.MakeEntryList(associates, &deptEntries, true)
			lines = helpers.SubstituteString(
				lines,
				newLines,
				associatesSubstitution)
			helpers.OverwriteWithLines(filepath, lines)
		}
	}
}

// GenerateTeamExecutiveList inserts generated lists of executives into team pages.
func GenerateTeamExecutiveList(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry) {

	firstYear, lastYear := helpers.GetYearRange()
	for year := firstYear; year <= lastYear; year++ {
		execs := []associate.Entry{}
		for _, entry := range (*entries)[year] {
			if entry.IsExecutive() {
				execs = append(execs, entry)
			}
		}

		filepath := helpers.RelativeFilePath(year, lastYear, "team")
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		lines := helpers.ReadContentLines(filepath)
		newLines := associate.MakeEntryList(associates, &execs, false)
		lines = helpers.SubstituteString(
			lines,
			newLines,
			executivesSubstitution)
		helpers.OverwriteWithLines(filepath, lines)
	}
}
