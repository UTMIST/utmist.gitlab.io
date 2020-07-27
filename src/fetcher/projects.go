package fetcher

import (
	"io/ioutil"
	"log"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

func fetchProjects() map[int][]project.Project {

	files, err := ioutil.ReadDir(project.ProjectsFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	projects := map[int][]project.Project{}

	for _, f := range files {
		project := project.LoadProject(f.Name())
		firstYear, lastYear := helpers.GetYearRange(project.Years)
		if err != nil {
			continue
		}

		for y := firstYear; y <= lastYear; y++ {
			projects[y] = append(projects[y], project)
		}
	}

	return projects
}
