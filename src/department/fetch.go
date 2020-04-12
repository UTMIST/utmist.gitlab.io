package department

const departmentsSheetRange = 2

// LoadDescs returns a map of department descriptions.
func LoadDescs(deptDescs *map[string]string, data []interface{}) {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < departmentsSheetRange; i++ {
		data = append(data, "")
	}

	// Map the department to its description.
	dept := data[0].(string)
	desc := data[1].(string)

	(*deptDescs)[dept] = desc
}
