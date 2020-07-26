package generator

import (
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsListSubstitution = "[//]: # events"

// GenerateEventList inserts generated lists of events into the events page.
func GenerateEventList(eventMap *map[int][]event.Event) {

	firstYear, lastYear := helpers.GetYearRange()
	for year := firstYear; year <= lastYear; year++ {
		events := (*eventMap)[year]

		filepath := helpers.RelativeFilePath(year, lastYear, "events")
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		lines := helpers.ReadContentLines(filepath)
		newLines := event.GenerateListPage(&events)
		lines = helpers.SubstituteString(
			lines,
			newLines,
			eventsListSubstitution)
		helpers.OverwriteWithLines(filepath, lines)
	}
}
