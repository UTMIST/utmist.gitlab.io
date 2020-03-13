package generator

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// File locations.
const configCopyFilename = "assets/config.yaml"
const configFilename = "config.yaml"
const teamCopyFilename = "assets/team.md"
const teamFilename = "content/team/list.md"
const discordBase = "https://discord.gg/"

// GeneratePages generates the content pages for Events, Associates, and Projects.
func GeneratePages(
	events []event.Event,
	associates []associate.Associate,
	projects []project.Project) {

	// Start with config copy.
	configLines := helpers.ReadContentLines(configCopyFilename)
	teamFileLines := helpers.ReadContentLines(teamCopyFilename)

	// Generate associate pages and department navbar links.
	associate.GenerateAssociatePages(associates)
	associate.GenerateNavbarDeptLinks(&configLines)
	associate.GenerateDeptList(&teamFileLines)
	associate.GenerateVPList(&teamFileLines, associates)

	// Generate event pages navbar links.
	event.GenerateEventPages(events)
	event.GenerateNavbarEventLinks(events, &configLines)

	// Insert discord link.
	insertDiscordLink(&teamFileLines)
	insertDiscordLink(&configLines)

	// Overwrite config.
	helpers.OverwriteWithLines(teamFilename, teamFileLines)
	helpers.OverwriteWithLines(configFilename, configLines)
}

func insertDiscordLink(lines *[]string) {
	discordLink, exists := os.LookupEnv("DISCORD_LINK")
	if !exists {
		return
	}

	for i := range *lines {
		(*lines)[i] = strings.Replace((*lines)[i], discordBase, fmt.Sprintf("%s%s", discordBase, discordLink), -1)
	}
}
