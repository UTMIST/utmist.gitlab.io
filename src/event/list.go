package event

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsFilePath = "./content/events.md"
const eventsBasePath = "./assets/events.md"
const tablePadder = "   "

// Generate the main events list page (events.md).
func generateEventList(events []Event, buildings *map[string]Building) {
	// Get header lines of events.md.
	lines := helpers.ReadFileBase(eventsBasePath, len(events), 11)

	// Add each event to the list.
	for i := 0; i < len(events); i++ {
		title := events[i].Title
		filename := helpers.StringToFileName(events[i].Title)
		dateStr := events[i].DateTime.Format(helpers.PrintDateLayout)

		location, room := events[i].getLocation(buildings)

		eventListing := fmt.Sprintf("|[%s](../%s)|%s|%s|%s|%s|%s|%s|%s|%s|",
			title,
			filename,
			tablePadder,
			dateStr[:len(dateStr)-6],
			tablePadder,
			dateStr[len(dateStr)-6:],
			tablePadder,
			location,
			tablePadder,
			room,
		)
		lines = append(lines, eventListing)
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}
