package event

import (
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// Paths for the event files.
const eventsDirPath = "./content/events/"

// Generate a page for event, including image, content, location, date/time.
func (e *Event) generatePage(buildings *map[string]Building, index int) {

	helpers.GenerateLog(fmt.Sprintf("%s", e.Title))

	// Format date and generate page header.
	lines := helpers.GenerateFrontMatter(
		e.Title,
		helpers.PadDateWithIndex(index),
		e.Title,
		[]string{"Event", e.Type})

	// If there's an image and/or summary, include them.
	if len(e.ImageLink) > 0 {
		displayLink := strings.Replace(e.ImageLink, "open?", "u/0/uc?", 1)

		imageLine := fmt.Sprintf("![%s](%s)", e.Title, displayLink)
		lines = append(lines, imageLine)
	}
	if len(e.Summary) > 0 {
		lines = append(lines, fmt.Sprintf("\n%s", e.Summary))
	}

	// Clean up the file and add footer with date/time and location.
	lines = append(lines, "")
	lines = append(lines, helpers.Breakline)
	lines = append(lines, "")
	printedDateStr := fmt.Sprintf("Date/Time: **%s.**",
		e.DateTime.Format(helpers.PrintDateTimeLayout))
	lines = append(lines, printedDateStr)
	if location, room := e.getLocation(buildings); len(location) > 0 {
		lines = append(lines, "")
		printedLocStr := fmt.Sprintf("Location: **%s%s.**",
			location,
			func() string {
				if len(room) == 0 {
					return ""
				}
				return fmt.Sprintf(" %s", room)
			}())
		lines = append(lines, printedLocStr)
	}

	// If there's a post-link, include it.
	if len(e.PostLink) > 0 {
		lines = append(lines, "")
		printedLocStr := fmt.Sprintf("Slides/Feedback: [%s](%s).",
			e.PostLink,
			e.PostLink)
		lines = append(lines, printedLocStr)
	}

	filename := fmt.Sprintf("./content/events/%s.md",
		helpers.StringToFileName(e.Title))
	helpers.OverwriteWithLines(filename, lines)
}

// GeneratePages generates events main page and each event's page.
func GeneratePages(events *[]Event, description string) {
	// Get list of UofT buildings.
	buildings, err := getUofTBuildingsList()
	if err != nil {
		log.Fatal(err)
	}

	helpers.GenerateGroupLog("event")
	os.Mkdir(eventsDirPath, os.ModePerm)

	// Generate events main page.
	generateEventList(events, &buildings, description)

	// Generate each event page.
	for i, event := range *events {
		event.generatePage(&buildings, len(*events)-i)
	}
}
