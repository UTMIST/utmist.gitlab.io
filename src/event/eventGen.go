package event

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
)

// Paths for the event files.
const eventsDirPath = "./content/events/"
const eventsFilePath = "./content/events.md"

// Discord invite link to add to config.yaml.
const discordBase = "https://discord.gg/"

// Generate a page for an event, including main image, content, location and date/time.
func generateEventPage(name string, event Event, buildings *map[string]Building) {

	logger.GenerateLog(fmt.Sprintf("%s", name))

	// Create and open and defer close file.
	filename := event.titleToFilename()
	f, err := os.Create(fmt.Sprintf("./content/events/%s.md", filename))
	if err != nil {
		logger.GenerateErrorLog(fmt.Sprintf("%s", name))
	}
	defer f.Close()

	// Format date and generate page header.
	dateStr := event.DateTime.Format(hugo.FileDateLayout)
	hugo.GeneratePageHeader(f, name, dateStr, event.Summary, []string{"Event", event.Type})

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
	fmt.Fprintln(f, hugo.Breakline)
	printedDateStr := fmt.Sprintf("Date/Time: **%s.**", event.DateTime.Format(hugo.PrintDateLayout))
	fmt.Fprintln(f, printedDateStr)
	if location := event.getLocation(buildings); len(location) > 0 {
		fmt.Fprintln(f, "")
		printedLocStr := fmt.Sprintf("Location: **%s.**", location)
		fmt.Fprintln(f, printedLocStr)
	}
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

// GenerateNavbarEventLinks generates event links for the navbar dropdown menu.
func GenerateNavbarEventLinks(events []Event) {

	// Open and defer close config_base.yaml as template.
	configBase, err := os.Open(hugo.ConfigBase)
	if err != nil {
		log.Fatal(err)
	}
	defer configBase.Close()

	discordLink, discordLinkExists := os.LookupEnv("DISCORD_LINK")

	// Read lines from config_base.
	lines := []string{}
	scanner := bufio.NewScanner(configBase)
	for scanner.Scan() {
		line := scanner.Text()
		if discordLinkExists && strings.Index(line, discordBase) > -1 {
			line = fmt.Sprintf("%s%s", line, discordLink)
		}
		lines = append(lines, line)
	}

	// Search for correct place in config_base you'd add event listings.
	navbarIndex := 0
	for navbarIndex < len(lines) {
		if lines[navbarIndex] == hugo.Navbar {
			break
		}
		navbarIndex++
	}
	navbarIndex += navbarShift

	// Store the lines that go before/after the events in config.yaml.
	preLines := []string{}
	for j := 0; j <= navbarIndex; j++ {
		preLines = append(preLines, lines[j])
	}
	postLines := []string{}
	for j := navbarIndex + 1; j < len(lines); j++ {
		postLines = append(postLines, lines[j])
	}

	// Add event link lines into config.yaml.
	eventLines := []string{}
	for i := maxNavbarEvents - 1; i >= 0; i-- {
		filename := events[i].titleToFilename()
		newEvent := []string{
			fmt.Sprintf("        - title: \"%s\"", events[i].Title),
			fmt.Sprintf("          url: /events/%s", filename),
		}
		eventLines = append(newEvent, eventLines...)
	}

	// Stitch config.yaml back together with preLines and postLines.
	lines = append(preLines, eventLines...)
	lines = append(lines, postLines...)

	// Overwrite the config.yaml file.
	configFile, err := os.Create(hugo.Config)
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

	// Header of main events page.
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
		eventsFile.WriteString(listItem + "\n")
	}
	eventsFile.Close()
}

// Number of lines to shift when identifying navbar entry in config_base.yaml.
const navbarShift = 2

// Dictating how many individual links appear on the navbar list.
const maxNavbarEvents = 3
