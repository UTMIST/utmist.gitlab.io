package generator

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// GenerateProjectLists generates projects lists for the projects page.
func GenerateProjectLists(projectMap *map[int][]project.Project) {
	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		projects := (*projectMap)[y]

		filenames := []string{"academic", "engineering", "projects"}
		for _, filename := range filenames {
			filepath := helpers.RelativeFilePath(y, lastYear, filename)
			if _, err := os.Stat(filepath); err != nil {
				log.Println(err)
				continue
			}

			yearToProjects := project.GroupByType(&projects)
			projectTypes := []string{"academic", "applied", "infrastructure"}
			for _, projType := range projectTypes {
				lines := helpers.ReadContentLines(filepath)
				yearTypeProjects := yearToProjects[projType]
				newLines := project.MakeList(&(yearTypeProjects))
				lines = helpers.SubstituteString(
					lines,
					newLines,
					fmt.Sprintf("[//]: # %s-projects", projType))
				helpers.OverwriteWithLines(filepath, lines)
			}
		}

	}
}
