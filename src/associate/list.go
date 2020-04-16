package associate

import (
	"sort"
)

// GenerateExecList generates a list of executive members.
func GenerateExecList(lines *[]string,
	associates *[]Associate, description string) {

	// Get list of execs.
	execs := []Associate{}
	for _, associate := range *associates {
		if associate.IsExec() && !associate.HasRetired() {
			execs = append(execs, associate)
		}
	}
	sort.Sort(List(execs))

	// Build the list of lines from the execs.
	newLines := []string{}
	if len(description) > 0 {
		newLines = append(newLines, description)
	}
	for _, exec := range execs {
		newLines = append(newLines, exec.GetEntry("", false, true))
	}
	newLines = append(newLines, "")
	(*lines) = append(*lines, newLines...)
}
