package associate

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// GroupByDept groups associates into their own department list.
func GroupByDept(associates *[]Associate) map[string][]Associate {

	// Populate an empty list for every department.
	deptAssocSet := map[string]map[string]Associate{}
	deptAssociates := map[string][]Associate{}
	for _, dept := range helpers.GetDeptNames(true) {
		deptAssocSet[dept] = map[string]Associate{}
		deptAssociates[dept] = []Associate{}
	}

	// Insert associates into their appropriate department, if it exists.
	for _, assoc := range *associates {
		dept := assoc.Department
		if _, exists := deptAssocSet[dept]; !exists {
			continue
		}

		existing, exists := deptAssocSet[dept][assoc.UofTEmail]
		if !exists {
			deptAssocSet[dept][assoc.UofTEmail] = assoc
			continue
		}

		existing.Position = fmt.Sprintf("%s, %s",
			existing.Position,
			assoc.Position)
		deptAssocSet[dept][assoc.UofTEmail] = existing
	}

	for _, dept := range helpers.GetDeptNames(true) {
		for _, assoc := range deptAssocSet[dept] {
			deptAssociates[dept] = append(deptAssociates[dept], assoc)
		}
	}

	return deptAssociates
}
