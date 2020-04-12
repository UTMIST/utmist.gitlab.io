package fetcher

const descriptionSheetRange = 2

func loadPageDescs(pageDescs *map[string]string, data []interface{}) {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < descriptionSheetRange; i++ {
		data = append(data, "")
	}

	// Map the department to its description.
	page := data[0].(string)
	desc := data[1].(string)

	(*pageDescs)[page] = desc
}
