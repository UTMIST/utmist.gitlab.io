package generator

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/department"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// File locations.
const aboutCopyFilename = "assets/about.md"
const aboutFilename = "content/about.md"
const configCopyFilename = "assets/config.yaml"
const configFilename = "config.yaml"

// GeneratePages generates pages for Associates/Events/Positions/Projects.
func GeneratePages(
	assocs *[]associate.Associate,
	deptDescs *map[string]string,
	events *[]event.Event,
	positions *[]position.Position,
	pastProjs *[]project.Project,
	projs *[]project.Project) {

	// Generate associate/event/project pages.
	department.GeneratePages(assocs, deptDescs, positions, projs, pastProjs)
	department.GenerateTeamPage(assocs, positions)
	event.GeneratePages(events)
	position.GeneratePage(positions)
	project.GeneratePages(projs, pastProjs)

	// Generate about page.
	GenerateAboutPage(positions)
}

// GenerateConfig generates the configuration file for Hugo site.
func GenerateConfig(events *[]event.Event, projs *[]project.Project) {
	helpers.GenerateLog("config")

	// Start with config copy.
	lines := helpers.ReadContentLines(configCopyFilename)

	// Generate associate/event/project navbar links.
	department.GenerateNavbarDeptLinks(&lines)
	event.GenerateNavbarEventLinks(events, &lines)
	project.GenerateNavbarProjectLinks(projs, &lines)

	// Insert discord link.
	helpers.InsertDiscordLink(&lines)

	// Overwrite config.
	helpers.OverwriteWithLines(configFilename, lines)
}

// GenerateAboutPage generates the about page.
func GenerateAboutPage(positions *[]position.Position) {
	helpers.GenerateLog("about")

	lines := helpers.ReadContentLines(aboutCopyFilename)
	lines = append(lines, helpers.GetJoinLines()...)

	helpers.OverwriteWithLines(aboutFilename, lines)
}
