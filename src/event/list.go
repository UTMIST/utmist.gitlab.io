package event

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
)

const eventsFilePath = "./content/events.md"

// Generate the main events list page (events.md).
func generateEventList(events []Event, buildings *map[string]Building) {
	// Get header lines of events.md.
	lines := readEventsFileBase()

	// Add each event into the list.
	for i := 0; i < len(events); i++ {
		title := events[i].Title
		filename := events[i].titleToFilename()
		dateStr := events[i].DateTime.Format(hugo.PrintDateLayout)

		listItem := fmt.Sprintf("|[%s](%s)|%s|%s|%s|",
			title,
			filename,
			dateStr[:len(dateStr)-6],
			dateStr[len(dateStr)-6:],
			events[i].getLocation(buildings),
		)
		lines = append(lines, listItem)
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}

// Reads the existing events.md and truncates it to the header.
func readEventsFileBase() []string {
	eventsFile, err := os.Open(eventsFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer eventsFile.Close()

	// Read lines from config_base.
	lines := []string{}
	scanner := bufio.NewScanner(eventsFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines[:11]
}
