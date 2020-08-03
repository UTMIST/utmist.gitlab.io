package project

import (
	"fmt"
)

const emptyListLine = "New projects coming soon! Or you can [propose one](https://github.com/UTMIST/Developer-Guide)!"

// MakeList creates a list of project lines.
func MakeList(projects *[]Project) []string {

	var lines []string
	if len(*projects) == 0 {
		return []string{emptyListLine}
	}

	for _, proj := range *projects {
		title := proj.Title
		if len(proj.External) > 0 {
			title = fmt.Sprintf("[%s](%s)", proj.Title, proj.External)
		}

		lines = append(lines, fmt.Sprintf("##### **%s**", title))
		if len(proj.Summary) > 0 {
			lines = append(lines, proj.Summary)
		}

		if len(proj.JoinInstructions) > 0 {
			lines = append(lines,
				fmt.Sprintf("- _Joining_: %s", proj.JoinInstructions))
		}

		lines = append(lines, "")
	}

	return lines
}
