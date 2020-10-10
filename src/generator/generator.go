package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/project"

	"gitlab.com/utmist/utmist.gitlab.io/src/bundle"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const substitutionPrefix = "[//]: # "

const associatesPattern = "associates"
const departmentsPattern = "departments"
const eventsPattern = "events"
const executivesPattern = "executives"
const insertPattern = "insert"
const positionsPattern = "positions"
const projectsPattern = "projects"
const seriesPattern = "series"
const webRawPattern = "webraw"
const yearsPattern = "years"
const projectMemberPattern = "team"

// InsertGeneratedSubstitutions inserts generated substitution lists.
func InsertGeneratedSubstitutions(bundle *bundle.Bundle, directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files { // iterate on every file in this directory
		if f.IsDir() { // if this file is a directory, make a recursive call
			InsertGeneratedSubstitutions(
				bundle,
				fmt.Sprintf("%s/%s", directory, f.Name()))
			continue
		}

		filepath := fmt.Sprintf("%s/%s", directory, f.Name())
		if !strings.HasSuffix(filepath, helpers.MarkdownExt) {
			continue // if this file is not markdown type, skip
		}

		oldLines := helpers.ReadContentLines(filepath)
		newLines := []string{}
		for lineIndex := 0; lineIndex < len(oldLines); lineIndex++ {
			line := oldLines[lineIndex]
			if len(line) < len(substitutionPrefix) ||
				line[:len(substitutionPrefix)] != substitutionPrefix {
				newLines = append(newLines, line)
				continue // if this line is not a title line, keep it
			}

			// iterate on every line in this file
			var theseLines []string

			pattern := strings.Split(line[len(substitutionPrefix):], "\\")
			year, err := strconv.Atoi(pattern[len(pattern)-1])
			if err != nil {
				year = -1
			}

			// match pattern[0] against insertion rules
			switch pattern[0] {

			case associatesPattern:
				if len(pattern) == 2 {
					theseLines = GenerateTeamAssociateList(
						bundle.Associates,
						bundle.Entries,
						year,
						false)
				} else if len(pattern) == 3 {
					theseLines = GenerateDepartmentAssociateLists(
						bundle.Associates,
						bundle.Entries,
						pattern[1],
						year)
				}

			case departmentsPattern: // insert department list
				theseLines = GenerateTeamDepartmentList(bundle.Entries, year)

			case eventsPattern: // insert event list
				theseLines = GenerateEventList(bundle.Events, year)

			case executivesPattern: // insert exec profiles
				theseLines = GenerateTeamAssociateList(
					bundle.Associates,
					bundle.Entries,
					year,
					true)

			case insertPattern: // insert local insertions
				theseLines = helpers.ReadContentLines(
					strings.Join(pattern[1:], "-"))

			case positionsPattern: // insert open positions
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

			case projectMemberPattern: // insert project members
				theseLines = GenerateProjectAssociateList(
					bundle.Associates,
					bundle.TeamEntries,
					strings.Join(pattern[1:], " "))

			case projectsPattern: // insert project list
				theseLines = GenerateProjectLists(
					bundle.Projects,
					pattern[1],
					year)

			case seriesPattern: // list other events in same series
				theseLines = []string{getSeriesString(
					directory,
					pattern[1],
					strings.Join(pattern[2:], "-"))}

			case webRawPattern: // insert web raw text resource
				theseLines = project.DownloadReadMe(
					strings.Join(pattern[1:], "-"))

			case yearsPattern: // insert year list
				theseLines = []string{getYearListString(
					helpers.StringToSimplePath(filepath), year)}

			}

			// add new lines
			newLines = append(newLines, theseLines...)
		}

		helpers.OverwriteWithLines(filepath, newLines)
	}
}
