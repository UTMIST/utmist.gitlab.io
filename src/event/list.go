package event

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsFilePath = "./content/events.md"

// GenerateListPage the main events list page (events.md).
func GenerateListPage(events *[]Event,
	buildings *map[string]Building, description string) {

	// Generate header for events.md
	lines := helpers.GenerateHeader("Events", "0001-01-04")
	if len(description) > 0 {
		lines = append(lines,
			description, "",
			helpers.Breakline)
	}

	// Add each event to the list.
	for _, event := range *events {
		lines = append(lines, event.insertListEntry(buildings, true)...)
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}
