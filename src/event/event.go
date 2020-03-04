package event

import (
	"strings"
	"time"
)

// Event represents an entry in the Events Google Sheet
type Event struct {
	Title    string
	Type     string
	DateTime time.Time
	Location string

	Summary   string
	ImageLink string
	PreLink   string
	PostLink  string
}

// List defines a list of events.
type List []Event

// Method Len() to implement sort.Sort.
func (e List) Len() int {
	return len(e)
}

// Method Less() to implement sort.Sort.
func (e List) Less(i, j int) bool {
	return e[j].DateTime.Before(e[i].DateTime)
}

// Method Swap() to implement sort.Sort.
func (e List) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// Determine the appropriate filename for this event.
func (e *Event) titleToFilename() string {
	// We use lowercase page paths.
	filename := strings.ToLower(strings.ToLower(e.Title))

	// Remove illegal characters from filenames.
	strsToRemove := []string{"'", ":", ",", "(", ")"}
	for _, strToRemove := range strsToRemove {
		filename = strings.Replace(filename, strToRemove, "", -1)
	}
	filename = strings.Replace(filename, " - ", " ", -1)
	filename = strings.Replace(filename, " ", "-", -1)

	return filename
}

// Parse location from event, for something familiar like a UofT building.
func (e *Event) getLocation(buildings *map[string]Building) (string, string) {

	// Definitely not building code.
	if len(e.Location) <= 2 {
		return e.Location, ""
	}

	// Try to find a UofT building code
	bldgCode := e.Location[:2]
	bldg, exists := (*buildings)[bldgCode]
	if exists && bldgCode == strings.ToUpper(bldgCode) {
		return bldg.getUofTMapsLink(e.Location), e.Location
	}

	return e.Location, ""
}
