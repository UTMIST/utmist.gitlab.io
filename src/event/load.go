package event

import (
	"fmt"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// EventsFolderPath describes where event page files are found.
const EventsFolderPath = "content/events/"

const dateTimePrefix = "datetime:"
const locationPrefix = "location:"
const titlePrefix = "title:"

// LoadEvent loads an event from a spreadsheet row.
func LoadEvent(filename string) Event {

	event := Event{}

	lines := helpers.ReadContentLines(
		fmt.Sprintf("%s%s", EventsFolderPath, filename))
	for _, line := range lines {
		if strings.Contains(line, dateTimePrefix) {
			event.DateTime = helpers.FormatDateEST(
				helpers.ColonRemainder(line))
		}
		if strings.Contains(line, locationPrefix) {
			event.Location = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, titlePrefix) {
			event.Title = strings.Trim(helpers.ColonRemainder(line), "\"")
		}
	}

	return event
}
