package event

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsSheetRange = 8

// LoadEvent loads an event from a spreadsheet row.
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
