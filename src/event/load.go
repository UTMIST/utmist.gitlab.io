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

	lines := helpers.ReadContentLines(fmt.Sprintf("%s%s", EventsFolderPath, filename))

	colonRemainder := func(line string) string {
		return strings.TrimSpace(line[strings.Index(line, ":")+1:])
	}

	for _, line := range lines {
		if strings.Contains(line, dateTimePrefix) {
			dateStr := colonRemainder(line)
			event.DateTime = helpers.FormatDateEST(dateStr)
		}
		if strings.Contains(line, locationPrefix) {
			event.Location = colonRemainder(line)
		}
		if strings.Contains(line, titlePrefix) {
			verboseTitle := colonRemainder(line)
			event.Title = strings.Trim(verboseTitle, "\"")
		}
	}

	return event
}
