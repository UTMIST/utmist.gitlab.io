package generator

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// File locations.
const configCopyFilename = "assets/config.yaml"
const configFilename = "config.yaml"
const teamCopyFilename = "assets/team.md"
const teamFilename = "content/team.md"
const projectCopyFilename = "assets/projects.md"
const projectFilename = "content/projects.md"
const discordBase = "https://discord.gg/"

// GeneratePages generates the content pages for Events, Associates, and Projects.
func GeneratePages(
	events []event.Event,
	associates []associate.Associate,
	positions []position.Position,
	projects []project.Project,
	pastProjects []project.Project) {

	// Start with config copy.
	configLines := helpers.ReadContentLines(configCopyFilename)
	teamListLines := helpers.ReadContentLines(teamCopyFilename)
	projectListLines := helpers.ReadContentLines(projectCopyFilename)

	// Generate associate pages and department navbar links.
	associate.GenerateAssociatePages(associates, positions)
	associate.GenerateNavbarDeptLinks(&configLines)
	associate.GenerateDeptList(&teamListLines)
	associate.GenerateVPList(&teamListLines, associates)

	// Load table of open positions into team list.
	teamListLines = append(teamListLines, position.MakeTable(positions)...)

	// Generate event pages and event navbar links.
	event.GenerateEventPages(events)
	event.GenerateNavbarEventLinks(events, &configLines)

	// Generate project navbar links.
	project.GenerateNavbarProjectLinks(projects, &configLines)

	// Load tables of active/past projects.
	projectListLines = append(projectListLines, project.MakeTable(projects, true)...)
	projectListLines = append(projectListLines, project.MakeTable(pastProjects, false)...)

	// Insert discord link.
	insertDiscordLink(&teamListLines)
	insertDiscordLink(&configLines)

	// Overwrite config.
	helpers.OverwriteWithLines(projectFilename, projectListLines)
	helpers.OverwriteWithLines(teamFilename, teamListLines)
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
