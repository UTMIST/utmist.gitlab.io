package helpers

import (
	"fmt"
)

// Breakline denotes the linebreak in a markdown file.
const Breakline = "---"

// OpenPositions denotes the header for the open position list.
const OpenPositions = "## **Open Positions**"

// PrintDateTimeLayout defines the layout we print out.
const PrintDateTimeLayout = "Monday, January 2, 2006, 3:04PM"

// PrintDateLayout defines the layout we print out.
const PrintDateLayout = "Monday, January 2, 2006"

const hideLastModified = "hideLastModified: true"
const includeFooter = "include_footer: true"
const sidebar = "sidebar: true"
const sidebarlogo = "sidebarlogo: whiteside"

// Format list of tags into a front matter string.
func getTagsListStr(tags []string) string {
	tagsStr := "tags: ["
	for _, tag := range tags {
		tagsStr = fmt.Sprintf("%s\"%s\",", tagsStr, tag)
	}
	if len(tags) > 0 {
		tagsStr = tagsStr[:len(tagsStr)-1]
	}

	tagsStr = fmt.Sprintf("%s]", tagsStr)

	return tagsStr
}

// GenerateHeader writes lines for a page header.
func GenerateHeader(title, date string) []string {

	header := []string{
		Breakline,
		fmt.Sprintf("title: \"%s\"", title),
		fmt.Sprintf("date: %s", date),
		hideLastModified,
		sidebar,
		sidebarlogo,
		includeFooter,
		Breakline,
		"",
	}
	return header
}

// StitchPageLink stitches new lines into the config.
func StitchPageLink(lines *[]string,
	objects []string, section, start string) {

	// Add new objects into config.
	newLines := []string{}
	for i := 0; i < len(objects); i++ {
		filename := StringToFileName(objects[i])
		newObj := []string{
			fmt.Sprintf("        - title: \"%s\"", objects[i]),
			fmt.Sprintf("          url: %s%s", section, filename),
		}

		newLines = append(newLines, newObj...)
	}

	StitchIntoLines(lines, &newLines, start, 1)
}

// StitchExternalLink stitches new lines into the config.
func StitchExternalLink(lines *[]string,
	titles, links []string, start string) {

	// Add new objects into config.
	newLines := []string{}
	for i := 0; i < len(titles); i++ {
		link := links[i]
		newObj := []string{
			fmt.Sprintf("        - title: \"%s\"", titles[i]),
			fmt.Sprintf("          url: %s", link),
		}

		newLines = append(newLines, newObj...)
	}

	StitchIntoLines(lines, &newLines, start, 1)
}
