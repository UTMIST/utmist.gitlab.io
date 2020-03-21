package position

const positionSheetRange = 5

const positionBasePath = "./assets/positions.md"
const positionDeptBasePath = "./assets/positions_dept.md"

// Position represents an open associate position at UTMIST.
type Position struct {
	Title        string
	Department   string
	Description  string
	Requirements string
	Instructions string
}

// LoadPosition loads a position from a spreadsheet row.
func LoadPosition(data []interface{}) Position {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < positionSheetRange; i++ {
		data = append(data, "")
	}

	position := Position{
		Title:        data[0].(string),
		Department:   data[1].(string),
		Description:  data[2].(string),
		Requirements: data[3].(string),
		Instructions: data[4].(string),
	}

	return position
}

// GroupByDept groups positions into their own department list.
func GroupByDept(positions *[]Position) map[string][]Position {
	deptPositions := map[string][]Position{}
	for _, pos := range *positions {
		posList, exists := deptPositions[pos.Department]
		if !exists {
			deptPositions[pos.Department] = []Position{}
		}
		deptPositions[pos.Department] = append(posList, pos)
	}
	return deptPositions
}
