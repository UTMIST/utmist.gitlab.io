package generator

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
)

const assocPositionsSubstitution = "[//]: # assoc-positions"
const execPositionsSubstitution = "[//]: # exec-positions"
const positionsSubstitution = "[//]: # positions"
const positionsPagePath = "content/recruitment.md"

// GeneratePositionList generates position lists for the recruitment page.
func GeneratePositionList(positions *[]position.Position) {
	execPositions := []position.Position{}
	assocPositions := []position.Position{}

	for _, pos := range *positions {
		if pos.IsExec() {
			execPositions = append(execPositions, pos)
			continue
		}
		assocPositions = append(assocPositions, pos)
	}

	lines := helpers.ReadContentLines(positionsPagePath)
	lines = helpers.SubstituteString(
		lines,
		position.MakeList(
			&execPositions,
			false,
			"Executive"),
		execPositionsSubstitution)
	lines = helpers.SubstituteString(
		lines,
		position.MakeList(
			&execPositions,
			false,
			"Associate"),
		assocPositionsSubstitution)

	helpers.OverwriteWithLines(positionsPagePath, lines)
}

// GenerateDeptPositionLists generates position lists for department pages.
func GenerateDeptPositionLists(positions *[]position.Position) {
	_, year := helpers.GetYearRange()
	depts := strings.Split(os.Getenv(fmt.Sprintf("DEPTS_%d", year)), ",")
	deptToPositions := map[string][]position.Position{}
	for _, dept := range depts {
		deptToPositions[dept] = []position.Position{}
	}

	for _, pos := range *positions {
		if _, exists := deptToPositions[pos.Department]; !exists {
			continue
		}
		deptToPositions[pos.Department] =
			append(deptToPositions[pos.Department], pos)
	}

	for dept, deptPositions := range deptToPositions {
		filepath := helpers.RelativeFilePath(year, year, dept)
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		sort.Sort(position.List(deptPositions))
		lines := helpers.ReadContentLines(filepath)

		lines = helpers.SubstituteString(
			lines,
			position.MakeList(
				&deptPositions,
				false,
				"Associate/Executive"),
			positionsSubstitution)

		helpers.OverwriteWithLines(filepath, lines)
	}
}
