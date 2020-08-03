package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/bundle"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const substitutionPrefix = "[//]: # "

const associatesPattern = "associates"
const departmentsPattern = "departments"
const eventsPattern = "events"
const executivesPattern = "executives"
const positionsPattern = "positions"
const projectsPattern = "projects"
const yearsPattern = "years"

// InsertGeneratedSubstitutions inserts generated substitution lists.
func InsertGeneratedSubstitutions(bundle *bundle.Bundle, directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			InsertGeneratedSubstitutions(
				bundle,
				fmt.Sprintf("%s/%s", directory, f.Name()))
			continue
		}

		filepath := fmt.Sprintf("%s/%s", directory, f.Name())
		if !strings.HasSuffix(filepath, helpers.MarkdownExt) {
			continue
		}

		oldLines := helpers.ReadContentLines(filepath)
		newLines := []string{}
		for lineIndex := 0; lineIndex < len(oldLines); lineIndex++ {

			var theseLines []string

			line := oldLines[lineIndex]
			if len(line) < len(substitutionPrefix) ||
				line[:len(substitutionPrefix)] != substitutionPrefix {
				newLines = append(newLines, line)
				continue
			}

			pattern := strings.Split(line[len(substitutionPrefix):], "-")
			switch len(pattern) {
			case 2:
				year, err := strconv.Atoi(pattern[1])
				if err != nil {
					year = -1
				}

				switch pattern[0] {
				case associatesPattern:
					theseLines = GenerateTeamAssociateList(
						bundle.Associates,
						bundle.Entries,
						year,
						false)
				case departmentsPattern:
					theseLines = GenerateTeamDepartmentList(bundle.Entries, year)
				case eventsPattern:
					theseLines = GenerateEventList(bundle.Events, year)
				case executivesPattern:
					theseLines = GenerateTeamAssociateList(
						bundle.Associates,
						bundle.Entries,
						year,
						true)
				case positionsPattern:
					positionsByDepts := (*bundle.PositionsByDepts)
					if _, exists := positionsByDepts[pattern[1]]; exists {
						theseLines = GenerateDeptPositionLists(
							bundle.PositionsByDepts,
							pattern[1])
					} else {
						theseLines = GeneratePositionList(
							bundle.PositionsByLevel,
							pattern[1])
					}
				}
			case 3:
				year, err := strconv.Atoi(pattern[2])
				if err != nil {
					newLines = append(newLines, line)
					continue
				}
				switch pattern[0] {
				case associatesPattern:
					theseLines = GenerateDepartmentAssociateLists(
						bundle.Associates,
						bundle.Entries,
						pattern[1],
						year)
				case projectsPattern:
					theseLines = GenerateProjectLists(
						bundle.Projects,
						pattern[1],
						year)
				case yearsPattern:
					theseLines = []string{getYearListString(
						helpers.StringToSimplePath(filepath), year)}
				}
			}
			newLines = append(newLines, theseLines...)
		}

		helpers.OverwriteWithLines(filepath, newLines)
	}
}
