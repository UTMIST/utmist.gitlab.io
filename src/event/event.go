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

func (e *Event) getLocationStr(bldgs *map[string]Building) string {

	location, room := e.getLocation(bldgs)
	if len(location) == 0 || len(room) == 0 {
		return ""
	}

	return fmt.Sprintf("- _Location_: %s %s", location, room)
}

func (e *Event) insertListEntry(bldgs *map[string]Building, list bool) []string {

	// Add lines for date/time, location, and breakline.
	newLines := []string{fmt.Sprintf("- _Date/Time_: **%s**",
		e.DateTime.Format(helpers.PrintDateTimeLayout)),
		e.getLocationStr(bldgs)}

	// If this is for a list, entry, return current lines and header.
	if list {
		return append([]string{
			// helpers.Breakline,
			fmt.Sprintf("##### **[%s](/events/%s)**",
				e.Title,
				helpers.StringToFileName(e.Title))}, newLines...)
	}

	// Insert links if they exist.
	if len(e.PreLink) > 0 {
		preStr := fmt.Sprintf("- _Signup/Preview_: [%s](%s).",
			helpers.GetURLBase(e.PreLink),
			e.PreLink)
		newLines = append(newLines, preStr)
	}
	if len(e.PostLink) > 0 {
		postStr := fmt.Sprintf("- _Slides/Feedback_: [%s](%s).",
			helpers.GetURLBase(e.PostLink),
			e.PostLink)
		newLines = append(newLines, postStr)
	}

	return newLines
}
