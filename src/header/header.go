package markdown

import (
	"fmt"
	"os"
)

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
		sidebar,
		sidebarlogo,
		Breakline,
		"",
	}

	for _, line := range header {
		fmt.Fprintln(f, line)
	}
}

// Breakline denotes the linebreak in a markdown file.
const Breakline = "---"
const sidebar = "sidebar: true"
const sidebarlogo = "sidebarlogo: whiteside"
