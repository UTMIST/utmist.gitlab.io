package event

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"

	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
)

// Paths for the event files.
const eventsDirPath = "./content/events/"

// Generate a page for an event, including main image, content, location and date/time.
func generateEventPage(name string, event Event, buildings *map[string]Building) {

	logger.GenerateLog(fmt.Sprintf("%s", name))

	// Format date and generate page header.
	dateStr := event.DateTime.Format(hugo.FileDateLayout)
	lines := hugo.GeneratePageHeader(name, dateStr, event.Summary, []string{"Event", event.Type})

	// If there's an image and/or summary, include them.
	if len(event.ImageLink) > 0 {
		displayLink := strings.Replace(event.ImageLink, "open?", "u/0/uc?", 1)

		imageLine := fmt.Sprintf("![%s](%s)", event.Title, displayLink)
		lines = append(lines, imageLine)
	}
	if len(event.Summary) > 0 {
		lines = append(lines, fmt.Sprintf("\n%s", event.Summary))
	}

	// Clean up the file and add footer with date/time and location.
	lines = append(lines, hugo.Breakline)
	printedDateStr := fmt.Sprintf("Date/Time: **%s.**", event.DateTime.Format(hugo.PrintDateLayout))
	lines = append(lines, printedDateStr)
	if location := event.getLocation(buildings); len(location) > 0 {
		lines = append(lines, "")
		printedLocStr := fmt.Sprintf("Location: **%s.**", location)
		lines = append(lines, printedLocStr)
	}

	filename := fmt.Sprintf("./content/events/%s.md", event.titleToFilename())
	helpers.OverwriteWithLines(filename, lines)
}

// GenerateEventPages generates events main page and each event's page.
func GenerateEventPages(events []Event) {
	// Get list of UofT buildings.
	buildings, err := getUofTBuildingsList()
	if err != nil {
		log.Fatal(err)
	}

	// Generate events main page.
	generateEventList(events, &buildings)

	// Create folder and generate each event page.
	os.Mkdir(eventsDirPath, os.ModePerm)
	logger.GenerateGroupLog("event")
	for _, event := range events {
		generateEventPage(event.Title, event, &buildings)
	}

}
