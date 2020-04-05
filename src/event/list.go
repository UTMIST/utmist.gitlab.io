package event

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsFilePath = "./content/events.md"
const eventsBasePath = "./assets/events.md"

// Generate the main events list page (events.md).
func generateEventList(events *[]Event, buildings *map[string]Building) {
	// Get header lines of events.md.
	lines := helpers.ReadFileBase(eventsBasePath, len(*events), 10)

	// Add each event to the list.
	for _, event := range *events {
		event.insertListEntry(&lines, buildings)
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}
