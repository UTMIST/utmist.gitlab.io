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
	firstYear, lastYear := helpers.GetYearRange(os.Getenv("YEARS"))
	for y := firstYear; y <= lastYear; y++ {
		events := (*eventMap)[y]

		filepath := helpers.RelativeFilePath(y, lastYear, "events")
		if _, err := os.Stat(filepath); err != nil {
			log.Println(err)
			continue
		}

		lines := helpers.ReadContentLines(filepath)
		eventLines := event.GenerateListPage(&events)
		lines = helpers.SubstituteString(
			lines,
			eventLines,
			eventsListSubstitution)

		yearLine := getYearListString("events", firstYear, lastYear, y)
		lines = helpers.SubstituteString(
			lines,
			[]string{yearLine},
			yearListSubstitution)
		helpers.OverwriteWithLines(filepath, lines)
	}
}
