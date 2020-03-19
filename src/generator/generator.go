package generator

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// File locations.
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
	associate.GenerateAssociatePages(associates, positions)
	associate.GenerateTeamPage(associates, positions)
	event.GenerateEventPages(events)
	project.GenerateProjectListPage(projects, pastProjects)
}

// GenerateConfig generates the configuration file for Hugo site.
func GenerateConfig(events *[]event.Event, projects *[]project.Project) {
	// Start with config copy.
	configLines := helpers.ReadContentLines(configCopyFilename)

	// Generate associate/event/project navbar links.
	associate.GenerateNavbarDeptLinks(&configLines)
	event.GenerateNavbarEventLinks(events, &configLines)
	project.GenerateNavbarProjectLinks(projects, &configLines)

	// Insert discord link.
	helpers.InsertDiscordLink(&configLines)

	// Overwrite config.
	helpers.OverwriteWithLines(configFilename, configLines)
}
