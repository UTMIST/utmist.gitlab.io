package event

import (
	"fmt"
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
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

func (e *Event) insertListEntry(lines *[]string, bldgs *map[string]Building) {
	title := e.Title
	filename := helpers.StringToFileName(e.Title)
	dateStr := e.DateTime.Format(helpers.PrintDateTimeLayout)

	// Add lines for event name and date/time.
	*lines = append(*lines, fmt.Sprintf("##### **[%s](%s)**", title, filename))
	*lines = append(*lines, fmt.Sprintf("- _Date/Time_: %s", dateStr))

	// If location is given, include a line for it.
	location, room := e.getLocation(bldgs)
	if len(location) > 0 {
		*lines = append(*lines, fmt.Sprintf("- _Location_: %s%s", location,
			func() string {
				if len(room) == 0 {
					return ""
				}
				return fmt.Sprintf(" %s", room)
			}()))
	}
	*lines = append(*lines, helpers.Breakline)
}
