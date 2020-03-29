package associate

import (
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const execListStart = "## **Leadership**"

// GenerateExecList generates a list of executive members.
func GenerateExecList(lines *[]string, associates *[]Associate) {

	// Add each exec to the list.
	execs := []Associate{}
	for _, associate := range *associates {
		if associate.IsExec() && !associate.HasGraduated() {
			execs = append(execs, associate)
		}
	}
	sort.Sort(List(execs))

	newLines := []string{}
	for _, exec := range execs {
		execLine := exec.GetLine("", false, true)
		newLines = append(newLines, execLine)
	}
	newLines = append(newLines, "")
	helpers.StitchIntoLines(lines, &newLines, execListStart, 1)
}
