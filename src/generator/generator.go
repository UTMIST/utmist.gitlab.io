package generator

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// GeneratePages generates the content pages for Events, Associates, and Projects.
func GeneratePages(events []event.Event, associates []associate.Associate, projects []project.Project) {
	associate.GenerateAssociatePages(associates)
	event.GenerateEventPages(events)
	event.GenerateNavbarEventLinks(events)
}
