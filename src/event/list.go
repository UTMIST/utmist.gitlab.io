package event

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsFilePath = "./content/events/list.md"
const eventsBasePath = "./assets/events.md"
const tablePadder = "   "
const blankDate = "date:"

// Generate the main events list page (events.md).
func generateEventList(events []Event, buildings *map[string]Building) {
	// Get header lines of events.md.
	lines := readEventsFileBase(len(events))

	// Add each event to the list.
	for i := 0; i < len(events); i++ {
		title := events[i].Title
		filename := helpers.StringToFileName(events[i].Title)
		dateStr := events[i].DateTime.Format(helpers.PrintDateLayout)

		location, room := events[i].getLocation(buildings)

		listItem := fmt.Sprintf("|[%s](../%s)|%s|%s|%s|%s|%s|%s|%s|%s|",
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
		lines = append(lines, listItem)
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}

// Reads the existing events.md and truncates it to the header.
func readEventsFileBase(num int) []string {

	lines := helpers.ReadContentLines(eventsBasePath)
	for i, line := range lines {
		if len(line) >= len(blankDate) && line[:len(blankDate)] == blankDate {
			line = fmt.Sprintf("%s %s", blankDate, helpers.PadDateWithIndex(num+1))
		}

		lines[i] = line
	}

	return lines[:11]
}
