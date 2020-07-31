package generator

import (
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const associateListSubstitution = "[//]: # associates"

// GenerateDepartmentAssociateLists inserts generated lists of associates into dept pages.
func GenerateDepartmentAssociateLists(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry) {

	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		depts := helpers.GetDeptNames(y)
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
