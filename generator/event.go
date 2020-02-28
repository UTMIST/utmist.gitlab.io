package generator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Paths for the event files.
const eventsDirPath = "./content/events/"
const eventsFilePath = "./content/events.md"

// Event represents an entry in the Events Google Sheet
type Event struct {
	Title    string
	Type     string
	DateTime time.Time
	Location string

	Summary   string
	ImageLink string
	PreLink   string
	PostLink  string
}

// Determine the appropriate filename for this event.
func (e *Event) titleToFilename() string {

	// Remove illegal characters from filenames.
	filename := strings.Replace(strings.ToLower(e.Title), "'", "", -1)
	filename = strings.Replace(filename, ":", "", -1)
	filename = strings.Replace(filename, ",", "", -1)
	filename = strings.Replace(filename, "(", "", -1)
	filename = strings.Replace(filename, ")", "", -1)
	filename = strings.Replace(filename, " - ", " ", -1)
	filename = strings.Replace(filename, " ", "-", -1)

	return filename
}

// Parse location from event, for something familiar like a UofT building.
func (e *Event) getLocation(buildings *map[string]Building) string {

	// Definitely not building code.
	if len(e.Location) <= 2 {
		return e.Location
	}

	// Try to find a UofT building code
	bldgCode := e.Location[:2]
	bldg, exists := (*buildings)[bldgCode]
	if exists && bldgCode == strings.ToUpper(bldgCode) {
		return bldg.getUofTMapsLink(e.Location)
	}

	return e.Location

}

// Generate a page for an event, including main image, content, location and date/time.
func generateEventPage(name string, event Event, buildings *map[string]Building) {

	generateLog(fmt.Sprintf("%s", name))

	// Create and open and defer close file.
	filename := event.titleToFilename()
	f, err := os.Create(fmt.Sprintf("./content/events/%s.md", filename))
	if err != nil {
		generateErrorLog(fmt.Sprintf("%s", name))
	}
	defer f.Close()

	// Format date and generate page header.
	dateStr := event.DateTime.Format(fileDateLayout)
	generatePageHeader(f, name, dateStr, event.Summary, []string{"Event", event.Type})

	// If there's an image and/or summary, include them.
	if len(event.ImageLink) > 0 {
		displayLink := strings.Replace(event.ImageLink, "open?", "u/0/uc?", 1)

		imageLine := fmt.Sprintf("![%s](%s)", event.Title, displayLink)
		fmt.Fprintln(f, imageLine)
	}
	if len(event.Summary) > 0 {
		fmt.Fprintln(f, fmt.Sprintf("\n%s", event.Summary))
	}

	// Clean up the file and add footer with date/time and location.
	fmt.Fprintln(f, breakLine)
	printedDateStr := fmt.Sprintf("Date/Time: **%s.**", event.DateTime.Format(printDateLayout))
	fmt.Fprintln(f, printedDateStr)
	if location := event.getLocation(buildings); len(location) > 0 {
		fmt.Fprintln(f, "")
		printedLocStr := fmt.Sprintf("Location: **%s.**", location)
		fmt.Fprintln(f, printedLocStr)
	}
}

// Generate events main page and each event's page.
func generateEventPages(events []Event) {
	// Get list of UofT buildings.
	buildings, err := getUofTBuildingsList()
	if err != nil {
		log.Fatal(err)
	}

	// Generate events main page.
	generateEventList(events, &buildings)

	// Create folder and generate each event page.
	os.Mkdir(eventsDirPath, os.ModePerm)
	generateGroupLog("event")
	for _, event := range events {
		generateEventPage(event.Title, event, &buildings)
	}

}

// GenerateNavbarEventLinks generates event links for the navbar dropdown menu.
func GenerateNavbarEventLinks(events []Event) {

	// Open and defer close config_base.yaml as template.
	configBase, err := os.Open(configBase)
	if err != nil {
		log.Fatal(err)
	}
	defer configBase.Close()

	// Read lines from config_base.
	lines := []string{}
	scanner := bufio.NewScanner(configBase)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Search for correct place in config_base you'd add event listings.
	i := 0
	for i < len(lines) {
		if lines[i] == navbar {
			break
		}
		i++
	}
	i += navbarShift

	// Store the lines that go before/after the events in config.yaml.
	preLines := []string{}
	postLines := []string{}
	for j := 0; j < len(lines); j++ {
		if j <= i {
			preLines = append(preLines, lines[j])
		} else {
			postLines = append(postLines, lines[j])
		}
	}

	// Add event link lines into config.yaml.
	eventLines := []string{}
	for i := len(events) - 1; i >= 0; i-- {
		filename := events[i].titleToFilename()
		newEvent := []string{
			fmt.Sprintf("        - title: \"%s\"", events[i].Title),
			fmt.Sprintf("          url: /events/%s", filename),
		}
		eventLines = append(newEvent, eventLines...)

		// Truncate the number of events shown on navbar.
		if len(eventLines) >= maxNavbarEvents {
			break
		}
	}

	// Stitch config.yaml back together with preLines and postLines.
	lines = append(preLines, eventLines...)
	lines = append(lines, postLines...)

	// Overwrite the config.yaml file.
	configFile, err := os.Create(config)
	if err != nil {
		log.Fatal(err)
	}
	configWrite := bufio.NewWriter(configFile)
	for _, line := range lines {
		configWrite.WriteString(line + "\n")
	}
	configWrite.Flush()
	configFile.Close()
}

// Generate the main events list page.
func generateEventList(events []Event, buildings *map[string]Building) {
	eventsFile, err := os.Create(eventsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range []string{
		"---",
		"title: Events",
		"date: 0001-01-04",
		"sidebar: true",
		"sidebarlogo: whiteside",
		"---",
		"We regularly host events, on our own or in collaboration with other organizations.\n",
		"|Event|Date|Time|Location|",
		"|-----|----|----|--------|",
	} {
		eventsFile.WriteString(line + "\n")
	}

	for i := 0; i < len(events); i++ {
		title := events[i].Title
		filename := events[i].titleToFilename()
		dateStr := events[i].DateTime.Format(printDateLayout)

		listItem := fmt.Sprintf("|[%s](%s)|%s|%s|%s|",
			title,
			filename,
			dateStr[:len(dateStr)-6],
			dateStr[len(dateStr)-6:],
			events[i].getLocation(buildings),
		)
		eventsFile.WriteString(listItem + "\n")
	}
	eventsFile.Close()
}
