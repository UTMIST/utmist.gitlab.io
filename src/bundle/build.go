package bundle

import (
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// BuildBundle maps years and departments to objects to create a bundle.
func BuildBundle(
	associates *map[string]associate.Associate,
	entries *map[int][]associate.Entry,
	events *map[int][]event.Event,
	positions *[]position.Position,
	projects *map[int][]project.Project) Bundle {

	bundle := Bundle{
		Associates:       associates,
		Entries:          buildEntryMap(entries),
		Events:           events,
		PositionsByDepts: buildPositionByDeptsMap(positions),
		PositionsByLevel: buildPositionByLevelMap(positions),
		Projects:         buildProjectMap(projects),
	}

	return bundle
}

func buildEntryMap(
	entries *map[int][]associate.Entry) *map[int]map[string][]associate.Entry {

	entryMap := map[int]map[string][]associate.Entry{}

	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		for _, entry := range (*entries)[y] {
			if _, exists := entryMap[y]; !exists {
				entryMap[y] = map[string][]associate.Entry{}
			}
			if _, exists := entryMap[y][entry.Department]; !exists {
				entryMap[y][entry.Department] = []associate.Entry{}
			}
			entryMap[y][entry.Department] =
				append(entryMap[y][entry.Department], entry)
		}
	}
	return &entryMap
}

func buildPositionByDeptsMap(
	positionList *[]position.Position) *map[string][]position.Position {

	positionMap := map[string][]position.Position{}
	for _, pos := range *positionList {
		if pos.IsExec() {
			positionMap["executive"] = append(positionMap["executive"], pos)
		} else {
			positionMap["associate"] = append(positionMap["executive"], pos)
		}
	}

	return &positionMap
}

func buildPositionByLevelMap(
	positionList *[]position.Position) *map[string][]position.Position {

	positionMap := map[string][]position.Position{}
	for _, pos := range *positionList {
		if _, exists := positionMap[pos.Department]; !exists {
			positionMap[pos.Department] = []position.Position{}
		}
		positionMap[pos.Department] =
			append(positionMap[pos.Department], pos)
	}

	return &positionMap
}

func buildProjectMap(
	projectList *map[int][]project.Project) *map[int]map[string][]project.Project {

	projectMap := map[int]map[string][]project.Project{}

	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		for _, proj := range (*projectList)[y] {
			if _, exists := projectMap[y]; !exists {
				projectMap[y] = map[string][]project.Project{}
			}
			if _, exists := projectMap[y][proj.Type]; !exists {
				projectMap[y][proj.Type] = []project.Project{}
			}
			projectMap[y][proj.Type] =
				append(projectMap[y][proj.Type], proj)
		}

	}

	return &projectMap
}
