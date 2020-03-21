package project

const projectSheetRange = 6

// ActiveStatus string for an active project status.
const ActiveStatus = "Active"

// Project represents an entry in the Projects Google Sheet
type Project struct {
	Title        string
	Status       string
	Department   string
	Description  string
	Link         string
	Instructions string
}

// LoadProject loads a project from a spreadsheet row.
func LoadProject(data []interface{}) Project {
	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < projectSheetRange; i++ {
		data = append(data, "")
	}

	project := Project{
		Title:        data[0].(string),
		Status:       data[1].(string),
		Department:   data[2].(string),
		Description:  data[3].(string),
		Link:         data[4].(string),
		Instructions: data[5].(string),
	}

	return project
}

// GroupByDept groups projects into their own department list.
func GroupByDept(projects *[]Project) map[string][]Project {
	deptProjects := map[string][]Project{}
	for _, proj := range *projects {
		projList, exists := deptProjects[proj.Department]
		if !exists {
			deptProjects[proj.Department] = []Project{}
		}
		deptProjects[proj.Department] = append(projList, proj)
	}
	return deptProjects
}
