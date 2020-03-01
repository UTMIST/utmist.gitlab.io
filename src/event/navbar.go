package event

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
)

// Discord invite link to add to config.yaml.
const discordBase = "https://discord.gg/"

// Number of lines to shift when identifying navbar entry in config_base.yaml.
const navbarShift = 2

// Dictating how many individual links appear on the navbar list.
const maxNavbarEvents = 3

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

	helpers.OverwriteWithLines(hugo.Config, lines)
}
