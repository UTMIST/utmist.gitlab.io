package bundle

import (
	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"
)

// Bundle is a struct containing the object maps for the website.
type Bundle struct {
	Associates       *map[string]associate.Associate
	Entries          *map[int]map[string][]associate.Entry
	Events           *map[int][]event.Event
	PositionsByDepts *map[string][]position.Position
	PositionsByLevel *map[string][]position.Position
	Projects         *map[int]map[string][]project.Project
}
