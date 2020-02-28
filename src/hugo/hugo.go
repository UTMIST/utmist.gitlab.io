package hugo

import (
	"fmt"
	"os"
)

// Breakline denotes the linebreak in a markdown file.
const Breakline = "---"

// Config defines the filename: config.yaml.
const Config = "config.yaml"

// ConfigBase defines the base filename: config_base.yaml.
const ConfigBase = "config_base.yaml"

// FileDateLayout defines the layout we write to files.
const FileDateLayout = "2006-01-02"

// MaxNavbarEvents defines how may events to show on the navbar.
const MaxNavbarEvents = 3

// Navbar is where the navbar entry in config_base.yaml begins.
const Navbar = "  navbar:"

// NavbarShift is number of lines to skip atthe main events page entry.
const NavbarShift = 2

// Sidebar is the markdown header property to show the sidebar.
const Sidebar = "sidebar: true"

// Sidebarlogo is the markdown header property to show the the whiteside logo.
const Sidebarlogo = "sidebarlogo: whiteside"

// PrintDateLayout defines the layout we print out.
const PrintDateLayout = "Mon, Jan 02 2006, 15:04"

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

// GeneratePageHeader writes a header to file given the page's front matter.
func GeneratePageHeader(f *os.File, title, date, summary string, tags []string) {
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

	for _, line := range header {
		fmt.Fprintln(f, line)
	}
}
