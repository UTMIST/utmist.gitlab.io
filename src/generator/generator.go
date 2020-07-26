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

// GenerateAssociateLists inserts generated lists of associates into pages.
func GenerateAssociateLists(
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
			deptNamePattern := helpers.StringToFileName(dept)
			filepath := fmt.Sprintf("content/%s-%d.md", deptNamePattern, year)
			if year == lastYear {
				filepath = fmt.Sprintf("content/%s.md", deptNamePattern)
			}

			if _, err := os.Stat(filepath); err != nil {
				log.Println(err)
				continue
			}

			lines := helpers.ReadContentLines(filepath)
			newLines := associate.MakeEntryList(associates, &deptEntries)
			lines = helpers.SubstituteString(
				lines,
				newLines,
				associatesSubstitution)
			helpers.OverwriteWithLines(filepath, lines)
		}
	}
}
