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
const aboutFilename = "content/about.md"
const aboutPhotoLink = "![AGM 2019 Photo](/images/intel.jpg)"
const configCopyFilename = "assets/config.yaml"
const configFilename = "config.yaml"

// GeneratePages generates pages for Associates/Events/Positions/Projects.
func GeneratePages(
	assocs *[]associate.Associate,
	descriptions *map[string]string,
	events *[]event.Event,
	positions *[]position.Position,
	pastProjs *[]project.Project,
	projs *[]project.Project,
	deptsFlag bool,
	eventsFlag bool) {

	// Generate associate/event/project pages.
	department.GenerateTeamPage(assocs, positions, descriptions)
	if deptsFlag {
		department.GeneratePages(assocs, descriptions, positions, projs, pastProjs)
	}
	if eventsFlag {
		event.GeneratePages(events, (*descriptions)["Events"])
	}
	position.GeneratePage(positions, descriptions)
	project.GeneratePages(projs, pastProjs, (*descriptions)["Project List"])

	// Generate about page.
	GenerateAboutPage(positions, descriptions)
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
func GenerateAboutPage(positions *[]position.Position,
	descriptions *map[string]string) {

	helpers.GenerateLog("about")

	lines := helpers.GenerateHeader("About Us", "0001-01-05")
	lines = append(lines, aboutPhotoLink, "", helpers.Breakline)
	if description, exists := (*descriptions)["About"]; exists {
		lines = append(lines, description, "", helpers.Breakline)
	}
	lines = append(lines, helpers.GetJoinLines((*descriptions)["Joining"])...)

	helpers.OverwriteWithLines(aboutFilename, lines)
}
