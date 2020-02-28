package generator

import (
	"fmt"
	"log"
	"os"
)

// Formats for Google Drive Sheets data to be formatted into.
const fileDateLayout = "2006-01-02"
const printDateLayout = "Mon, Jan 02 2006, 15:04"

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

func generatePageHeader(f *os.File, title, date, summary string, tags []string) {
	header := []string{
		breakLine,
		fmt.Sprintf("title: \"%s\"", title),
		fmt.Sprintf("date: %s", date),
		fmt.Sprintf("summary: \"%s\"", summary),
		getTagsListStr(tags),
		"hideLastModified: true",
		sidebar,
		sidebarlogo,
		breakLine,
		"",
	}

	for _, line := range header {
		fmt.Fprintln(f, line)
	}
}

func generateProjectPage() {

}

func generateProjectPages() {

}

// GeneratePages generates the content pages for Events, Execs, and Projects.
func GeneratePages(events []Event, execs []Exec, projects []Project) {
	generateExecPages(execs)
	generateEventPages(events)
}

func generateLog(str string) {
	log.Println(fmt.Sprintf("\tGenerating page for %s.", str))
}

func generateErrorLog(str string) {
	log.Println(fmt.Sprintf("\tFailed to generate page for %s.", str))
}

func generateGroupLog(str string) {
	log.Println(fmt.Sprintf("Generating %s pages.", str))
}

// generator uses config_base.yaml to insert the links we want in config.yaml.
const config = "config.yaml"
const configBase = "config_base.yaml"

// Identifying where the navbar entry in config_base.yaml begins.
const navbar = "  navbar:"

// Number of lines to shift when identifying navbar entry in config_base.yaml.
const navbarShift = 2

// Dictating how many individual links appear on the navbar list.
const maxNavbarEvents = 3

const breakLine = "---"
const sidebar = "sidebar: true"
const sidebarlogo = "sidebarlogo: whiteside"
