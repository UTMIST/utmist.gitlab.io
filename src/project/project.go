package project

// ActiveStatus string for an active project status.
const ActiveStatus = "Active"

// Project represents an entry in the Projects Google Sheet
type Project struct {
	Link             string
	Image            string
	JoinInstructions string
	Summary          string
	Title            string
	Type             string
	Years            string
}

// GroupByType groups projects into their own department list.
func GroupByType(projects *[]Project) map[string][]Project {
	projTypeMap := map[string][]Project{}
	for _, proj := range *projects {

		projList, exists := projTypeMap[proj.Type]
		if !exists {
			projTypeMap[proj.Type] = []Project{}
		}
		projTypeMap[proj.Type] = append(projList, proj)

	}
	return projTypeMap
}
