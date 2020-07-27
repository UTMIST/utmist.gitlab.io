package position

import (
	"strings"
	"time"
)

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
