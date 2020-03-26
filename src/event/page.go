package event

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"

	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
)

// Paths for the event files.
const eventsDirPath = "./content/events/"

// Generate a page for an event, including main image, content, location and date/time.
func generateEventPage(name string, event Event, buildings *map[string]Building, index int) {

	logger.GenerateLog(fmt.Sprintf("%s", name))

	// Format date and generate page header.
	lines := helpers.GenerateHugoPageHeader(name, helpers.PadDateWithIndex(index), event.Title, []string{"Event", event.Type})

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
	lines = append(lines, "")
	lines = append(lines, helpers.Breakline)
	lines = append(lines, "")
	printedDateStr := fmt.Sprintf("Date/Time: **%s.**", event.DateTime.Format(helpers.PrintDateLayout))
	lines = append(lines, printedDateStr)
	if location, room := event.getLocation(buildings); len(location) > 0 {
		lines = append(lines, "")
		printedLocStr := fmt.Sprintf("Location: **%s%s.**", location, func() string {
			if len(room) == 0 {
				return ""
			}
			return fmt.Sprintf(" %s", room)
		}())
		lines = append(lines, printedLocStr)
	}

	// If there's a post-link, include it.
	if len(event.PostLink) > 0 {
		lines = append(lines, "")
		printedLocStr := fmt.Sprintf("Slides/Feedback: [%s](%s)", event.PostLink, event.PostLink)
		lines = append(lines, printedLocStr)
	}

	filename := fmt.Sprintf("./content/events/%s.md", helpers.StringToFileName(event.Title))
	helpers.OverwriteWithLines(filename, lines)
}

// GenerateEventPages generates events main page and each event's page.
func GenerateEventPages(events *[]Event) {
	// Get list of UofT buildings.
	buildings, err := getUofTBuildingsList()
	if err != nil {
		log.Fatal(err)
	}

	logger.GenerateGroupLog("event")
	os.Mkdir(eventsDirPath, os.ModePerm)

	// Generate events main page.
	generateEventList(events, &buildings)

	// Generate each event page.
	for i, event := range *events {
		generateEventPage(event.Title, event, &buildings, len(*events)-i)
	}
}
