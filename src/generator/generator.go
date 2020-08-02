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
const contentDirectory = "./content/"

const associatesPattern = "associates"
const departmentsPattern = "departments"
const eventsPattern = "events"
const executivesPattern = "executives"
const positionsPattern = "positions"
const projectsPattern = "projects"
const yearsPattern = "years"

// InsertGeneratedSubstitutions inserts generated substitution lists.
func InsertGeneratedSubstitutions(bundle bundle.Bundle) {
	files, err := ioutil.ReadDir(contentDirectory)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((*bundle.Projects)[2019]["Academic"])

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filepath := fmt.Sprintf("%s%s", contentDirectory, f.Name())
		lines := helpers.ReadContentLines(filepath)
		lines2 := []string{}
		for lineIndex := 0; lineIndex < len(lines); lineIndex++ {

			var newLines []string

			line := lines[lineIndex]
			if len(line) < len(substitutionPrefix) ||
				line[:len(substitutionPrefix)] != substitutionPrefix {
				lines2 = append(lines2, line)
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
				case departmentsPattern:
					newLines = GenerateTeamDepartmentList(bundle.Entries, year)
				case eventsPattern:
					newLines = GenerateEventList(bundle.Events, year)
				case executivesPattern:
					newLines = GenerateTeamExecutiveList(
						bundle.Associates,
						bundle.Entries,
						year)
				case positionsPattern:
					positionsByDepts := (*bundle.PositionsByDepts)
					if _, exists := positionsByDepts[pattern[1]]; exists {
						newLines = GenerateDeptPositionLists(
							bundle.PositionsByDepts,
							pattern[1])
					} else {
						newLines = GeneratePositionList(
							bundle.PositionsByLevel,
							pattern[1])
					}
				}
			case 3:
				year, err := strconv.Atoi(pattern[2])
				if err != nil {
					lines2 = append(lines2, line)
					continue
				}
				switch pattern[0] {
				case associatesPattern:
					newLines = GenerateDepartmentAssociateLists(
						bundle.Associates,
						bundle.Entries,
						pattern[1],
						year)
				case projectsPattern:
					newLines = GenerateProjectLists(
						bundle.Projects,
						pattern[1],
						year)
				case yearsPattern:
					newLines = []string{getYearListString(
						helpers.StringToSimplePath(pattern[1]), year)}
				}
			}
			lines2 = append(lines2, newLines...)
		}

		helpers.OverwriteWithLines(filepath, lines2)
	}
}
