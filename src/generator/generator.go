package generator

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/exec"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// GeneratePages generates the content pages for Events, Execs, and Projects.
func GeneratePages(events []event.Event, execs []exec.Exec, projects []project.Project) {
	exec.GenerateExecPages(execs)
	event.GenerateEventPages(events)
	event.GenerateNavbarEventLinks(events)
}
