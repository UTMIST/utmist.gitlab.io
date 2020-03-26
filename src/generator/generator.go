package generator

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// File locations.
const aboutCopyFilename = "assets/about.md"
const aboutFilename = "content/about.md"
const configCopyFilename = "assets/config.yaml"
const configFilename = "config.yaml"

// GeneratePages generates the content pages for Events, Associates, and Projects.
func GeneratePages(
	events *[]event.Event,
	associates *[]associate.Associate,
	positions *[]position.Position,
	projects *[]project.Project,
	pastProjects *[]project.Project) {

	// Generate associate/event/project pages.
	associate.GenerateAssociatePages(associates, positions, projects, pastProjects)
	associate.GenerateTeamPage(associates, positions)
	event.GenerateEventPages(events)
	project.GenerateProjectListPage(projects, pastProjects)

	// Generate about page.
	GenerateAboutPage(positions)
}

// GenerateConfig generates the configuration file for Hugo site.
func GenerateConfig(events *[]event.Event, projects *[]project.Project) {
	logger.GenerateLog("config")

	// Start with config copy.
	lines := helpers.ReadContentLines(configCopyFilename)

	// Generate associate/event/project navbar links.
	associate.GenerateNavbarDeptLinks(&lines)
	event.GenerateNavbarEventLinks(events, &lines)
	project.GenerateNavbarProjectLinks(projects, &lines)

	// Insert discord link.
	helpers.InsertDiscordLink(&lines)

	// Overwrite config.
	helpers.OverwriteWithLines(configFilename, lines)
}

// GenerateAboutPage generates the about page.
func GenerateAboutPage(positions *[]position.Position) {
	logger.GenerateLog("about")

	lines := helpers.ReadContentLines(aboutCopyFilename)
	lines = append(lines, position.MakeList(positions, false)...)

	helpers.InsertDiscordLink(&lines)
	helpers.OverwriteWithLines(aboutFilename, lines)
}
