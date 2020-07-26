package event

import (
	"log"
)

const eventsFilePath = "content/events.md"

// GenerateListPage the main events list page (events.md).
func GenerateListPage(events *[]Event) []string {

	// Get list of UofT buildings.
	buildings, err := getUofTBuildingsList()
	if err != nil {
		log.Fatal(err)
	}

	// Add each event's listing lines to the list.
	lines := []string{}
	for _, event := range *events {
		lines = append(lines, event.insertListEntry(&buildings, true)...)
	}

	return lines
}
