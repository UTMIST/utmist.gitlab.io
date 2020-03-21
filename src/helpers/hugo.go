package helpers

import (
	"fmt"
)

// BlankDate is the dateless header line.
const BlankDate = "date:"

// Breakline denotes the linebreak in a markdown file.
const Breakline = "---"

// FileDateLayout defines the layout we write to files.
const FileDateLayout = "2006-01-02"

// MaxNavbarEvents defines how may events to show on the navbar.
const MaxNavbarEvents = 3

// OpenPositions denotes the header for the open position list.
const OpenPositions = "## **Open Positions**"

// PrintDateLayout defines the layout we print out.
const PrintDateLayout = "Mon, Jan 02 2006, 15:04"

// Sidebar is the markdown header property to show the sidebar.
const Sidebar = "sidebar: true"

// Sidebarlogo is the markdown header property to show the the whiteside logo.
const Sidebarlogo = "sidebarlogo: whiteside"

// TablePadder is the ghost column content for displaying tables.
const TablePadder = " "

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

// GenerateHugoPageHeader writes lines for a page header.
func GenerateHugoPageHeader(title, date, summary string, tags []string) []string {
	header := []string{
		Breakline,
		fmt.Sprintf("title: \"%s\"", title),
		fmt.Sprintf("date: %s", date),
		fmt.Sprintf("summary: \"%s\"", summary),
		getTagsListStr(tags),
		"hideLastModified: true",
		Sidebar,
		Sidebarlogo,
		Breakline,
		"",
	}
	return header
}

// StitchPageLink stitches new lines into the config.
func StitchPageLink(lines *[]string, objects []string, section, start string) {

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
func StitchExternalLink(lines *[]string, titles, links []string, start string) {

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

// TablePad returns a string with <num> padders.
func TablePad(num int) string {
	padding := ""
	for i := 0; i < num; i++ {
		padding += TablePadder
	}
	return padding
}
