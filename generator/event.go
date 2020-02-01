package generator

import (
	"strings"
	"time"
)

type Event struct {
	Title    string
	Type     string
	DateTime time.Time

	Summary   string
	ImageLink string
	PreLink   string
	PostLink  string
}

func (e *Event) titleToFilename() string {
	filename := strings.Replace(strings.ToLower(e.Title), "'", "", -1)
	filename = strings.Replace(filename, ":", "", -1)
	filename = strings.Replace(filename, ",", "", -1)
	filename = strings.Replace(filename, " - ", " ", -1)
	filename = strings.Replace(filename, " ", "-", -1)

	return filename
}
