package event

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
)

const eventsFilePath = "./content/events/list.md"
const eventsBasePath = "./assets/events.md"
const tablePadder = "   "
const blankDate = "date:"

// Generate the main events list page (events.md).
func generateEventList(events []Event, buildings *map[string]Building) {
	// Get header lines of events.md.
	lines := readEventsFileBase(len(events))

	// Add each event into the list.
	for i := 0; i < len(events); i++ {
		title := events[i].Title
		filename := events[i].titleToFilename()
		dateStr := events[i].DateTime.Format(hugo.PrintDateLayout)

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
	eventsFile, err := os.Open(eventsBasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer eventsFile.Close()

	// Read lines from config_base.
	lines := []string{}
	scanner := bufio.NewScanner(eventsFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) >= len(blankDate) && line[:len(blankDate)] == blankDate {
			line = fmt.Sprintf("%s %s", blankDate, helpers.PadDateWithIndex(num+1))
		}
		lines = append(lines, line)
	}

	return lines[:11]
}
