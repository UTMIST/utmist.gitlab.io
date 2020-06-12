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

// GeneratePage includes image, content, location, date/time for an event.
func (e *Event) GeneratePage(buildings *map[string]Building, index int) {

	helpers.GenerateLog(fmt.Sprintf("%s", e.Title))

	// Format date and generate page header.
	lines := helpers.GenerateHeader(e.Title,
		helpers.PadDateWithIndex(index))

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
	lines = append(lines, "", helpers.Breakline, "")
	lines = append(lines, e.insertListEntry(buildings, false)...)

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
	GenerateListPage(events, &buildings, description)

	// Generate each event page.
	for i, event := range *events {
		event.GeneratePage(&buildings, len(*events)-i)
	}
}
