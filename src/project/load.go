package project

import (
	"fmt"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

// ProjectsFolderPath describes where event page files are found.
const ProjectsFolderPath = "content/projects/"

const linkPrefix = "link:"
const imagePrefix = "image:"
const joinPrefix = "join:"
const summaryPrefix = "summary:"
const titlePrefix = "title:"
const typePrefix = "type:"
const yearsPrefix = "years:"

// LoadProject loads a project from a file
func LoadProject(filename string) Project {

	filepath := fmt.Sprintf(
		"%s%s/%s",
		ProjectsFolderPath,
		filename,
		helpers.PageIndex)
	lines := helpers.ReadContentLines(filepath)

	project := Project{}
	for _, line := range lines {
		if strings.Contains(line, linkPrefix) {
			project.Link = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, imagePrefix) {
			project.Image = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, joinPrefix) {
			project.JoinInstructions = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, summaryPrefix) {
			project.Summary = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, titlePrefix) {
			project.Title = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, typePrefix) {
			project.Type = helpers.ColonRemainder(line)
		}
		if strings.Contains(line, yearsPrefix) {
			project.Years = helpers.ColonRemainder(line)
		}
	}

	if len(project.Link) <= 0 {
		project.Link = fmt.Sprintf("/projects/%s", filename)
	}

	project.Link = helpers.StringToSimplePath(project.Link)

	return project
}
