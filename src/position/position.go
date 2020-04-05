package position

import (
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const positionSheetRange = 6

// Position represents an open associate position at UTMIST.
type Position struct {
	Title        string
	Department   string
	Description  string
	Requirements string
	Instructions string
	Deadline     time.Time
}

// List defines a list of events.
type List []Position

// Method Len() to implement sort.Sort.
func (p List) Len() int {
	return len(p)
}

// Method Less() to implement sort.Sort.
func (p List) Less(i, j int) bool {
	if p[j].Deadline.After(p[i].Deadline) {
		return true
	} else if p[i].Deadline.After(p[j].Deadline) {
		return false
	}

	if strings.Compare(p[i].Title, p[j].Title) < 0 {
		return true
	}

	return false
}

// Method Swap() to implement sort.Sort.
func (p List) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// IsExec returns whether this position is executive.
func (p *Position) IsExec() bool {
	return strings.Index(p.Title, "VP") >= 0 ||
		strings.Index(p.Title, "President") >= 0
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

	if len(data[5].(string)) > 0 {
		position.Deadline = helpers.FormatDateEST(data[5].(string))
	}

	return position
}

// GroupByDept groups positions into their own department list.
func GroupByDept(positions *[]Position) map[string][]Position {
	deptPositions := map[string][]Position{}
	for _, dept := range helpers.GetDeptNames(false) {
		deptPositions[dept] = []Position{}
	}

	for _, pos := range *positions {
		posList, exists := deptPositions[pos.Department]
		if !exists {
			deptPositions[pos.Department] = []Position{}
		}
		deptPositions[pos.Department] = append(posList, pos)
	}
	return deptPositions
}
