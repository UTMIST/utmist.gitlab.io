package associate

import (
	"fmt"
	"sort"
)

// GenerateExecList generates a list of executive members.
func GenerateExecList(lines *[]string,
	associates *[]Associate, description string) {

	// Build the list of lines from the execs.
	newLines := []string{}
	if len(description) > 0 {
		newLines = append(newLines, description)
	}

	// Get set of unique execs.
	execSet := map[string]Associate{}
	for _, assoc := range *associates {
		if !assoc.IsExec(true) || assoc.HasRetired() {
			continue
		}

		if exec, exists := execSet[assoc.UofTEmail]; !exists {
			execSet[assoc.UofTEmail] = assoc
		} else {
			exec.Position = fmt.Sprintf("%s, %s",
				exec.Position,
				assoc.Position)
		}
	}

	// Load set into sorted list, and into lines.
	execs := []Associate{}
	for _, exec := range execSet {
		execs = append(execs, exec)
	}
	sort.Sort(List(execs))
	for _, exec := range execs {
		newLines = append(newLines, exec.GetEntry("", false, true))
	}
	newLines = append(newLines, "")
	(*lines) = append(*lines, newLines...)
}
