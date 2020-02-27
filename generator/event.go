package generator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

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

func (e *Event) titleToFilename() string {
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

func generateEventPage(name string, event Event, buildings *map[string]Building) {
	generateLog(fmt.Sprintf("%s", name))

	filename := event.titleToFilename()
	f, err := os.Create(fmt.Sprintf("./content/events/%s.md", filename))
	if err != nil {
		generateErrorLog(fmt.Sprintf("%s", name))
	}
	defer f.Close()

	dateStr := event.DateTime.Format(fileDateLayout)
	generatePageHeader(f, name, dateStr, event.Summary, []string{"Event", event.Type})

	if len(event.ImageLink) > 0 {
		displayLink := strings.Replace(event.ImageLink, "open?", "u/0/uc?", 1)

		imageLine := fmt.Sprintf("![%s](%s)", event.Title, displayLink)
		fmt.Fprintln(f, imageLine)
	}

	if len(event.Summary) > 0 {
		fmt.Fprintln(f, fmt.Sprintf("\n%s", event.Summary))
	}

	fmt.Fprintln(f, breakLine)

	printedDateStr := fmt.Sprintf("Date/Time: **%s.**", event.DateTime.Format(printDateLayout))
	fmt.Fprintln(f, printedDateStr)

	if location := event.getLocation(buildings); len(location) > 0 {
		fmt.Fprintln(f, "")
		printedLocStr := fmt.Sprintf("Location: **%s.**", location)
		fmt.Fprintln(f, printedLocStr)
	}
}

func generateEventPages(events []Event) {
	buildings, err := getUofTBuildingsList()
	if err != nil {
		log.Fatal(err)
	}

	generateEventList(events, &buildings)
	generateGroupLog("event")
	for _, event := range events {
		generateEventPage(event.Title, event, &buildings)
	}
}

// GenerateEventLinks generates event links for the navbar dropdown meny.
func GenerateEventLinks(events []Event) {
	configFile, err := os.Open(configBase)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()

	lines := []string{}
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] != navbar {
			i++
			continue
		}
		i += navbarShift

		preLines := []string{}
		postLines := []string{}
		for j := 0; j < len(lines); j++ {
			if j <= i {
				preLines = append(preLines, lines[j])
			} else {
				postLines = append(postLines, lines[j])
			}
		}

		eventLines := []string{}

		for i := len(events) - 1; i >= 0; i-- {
			filename := events[i].titleToFilename()
			newEvent := []string{
				fmt.Sprintf("        - title: \"%s\"", events[i].Title),
				fmt.Sprintf("          url: /events/%s", filename),
			}

			eventLines = append(newEvent, eventLines...)
		}

		if len(eventLines) > maxNavbarEvents*2 {
			eventLines = eventLines[:maxNavbarEvents*2]
		}

		lines = append(preLines, eventLines...)
		lines = append(lines, postLines...)

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

		break
	}
}

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
