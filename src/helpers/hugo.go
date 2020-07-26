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
const sidebar = "sidebar: false"
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
