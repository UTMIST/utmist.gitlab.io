package event

import (
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsSheetRange = 8

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
func (e *Event) getLocation(buildings *map[string]Building) string {

	// Definitely not building code.
	if len(e.Location) <= 2 {
		return e.Location
	}

	// Try to find a UofT building code
	bldgCode := e.Location[:2]
	bldg, exists := (*buildings)[bldgCode]
	if exists && bldgCode == strings.ToUpper(bldgCode) {
		return bldg.getUofTMapsLink(e.Location)
	}

	return e.Location
}

// Load an event from a spreadsheet row.
func LoadEvent(data []interface{}) Event {
	for i := len(data); i < eventsSheetRange; i++ {
		data = append(data, "")
	}

	event := Event{
		Title:     data[0].(string),
		Type:      data[1].(string),
		DateTime:  helpers.FormatDateEST(data[2].(string)),
		Location:  data[3].(string),
		Summary:   data[4].(string),
		ImageLink: data[5].(string),
		PreLink:   data[6].(string),
		PostLink:  data[7].(string),
	}

	return event
}
