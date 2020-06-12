package project

import "gitlab.com/utmist/utmist.gitlab.io/src/helpers"

// ActiveStatus string for an active project status.
const ActiveStatus = "Active"

// Project represents an entry in the Projects Google Sheet
type Project struct {
	Title        string
	Status       string
	Department   string
	Description  string
	Instructions string
	Link         string
}

// GroupByDept groups projects into their own department list.
func GroupByDept(projects *[]Project) map[string][]Project {
	deptProjects := map[string][]Project{}
	for _, dept := range helpers.GetDeptNames(false) {
		deptProjects[dept] = []Project{}
	}

	for _, proj := range *projects {
		projList, exists := deptProjects[proj.Department]
		if !exists {
			deptProjects[proj.Department] = []Project{}
		}
		deptProjects[proj.Department] = append(projList, proj)
	}
	return deptProjects
}
