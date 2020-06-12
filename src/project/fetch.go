package project

const projectSheetRange = 6

// Load loads a project from a spreadsheet row.
func Load(data []interface{}) Project {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < projectSheetRange; i++ {
		data = append(data, "")
	}

	project := Project{
		Title:        data[0].(string),
		Status:       data[1].(string),
		Department:   data[2].(string),
		Description:  data[3].(string),
		Instructions: data[4].(string),
		Link:         data[5].(string),
	}

	return project
}
