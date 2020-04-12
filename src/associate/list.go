package associate

import (
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const execListStart = "### **Leadership**"

// GenerateExecList generates a list of executive members.
func GenerateExecList(lines *[]string,
	associates *[]Associate, description string) {

	// Get list of execs.
	execs := []Associate{}
	for _, associate := range *associates {
		if associate.IsExec() && !associate.HasGraduated() {
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
		execLine := exec.GetLine("", false, true)
		newLines = append(newLines, execLine)
	}
	newLines = append(newLines, "")
	helpers.StitchIntoLines(lines, &newLines, execListStart, 1)
}
