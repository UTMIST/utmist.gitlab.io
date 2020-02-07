package generator

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
		breakLine,
		"",
	}

	for _, line := range header {
		fmt.Fprintln(f, line)
	}
}

func generateEventPage(name string, event Event) {
	generateLog(fmt.Sprintf("%s", name))

	filename := event.titleToFilename()
	f, err := os.Create(fmt.Sprintf("./content/events/%s.md", filename))
	if err != nil {
		generateErrorLog(fmt.Sprintf("%s", name))
	}
	defer f.Close()

	dateStr := event.DateTime.Format(fileDateLayout)
	generatePageHeader(f, name, dateStr, event.Summary, []string{"Event", event.Type})

	if len(event.ImageLink) > 0 {
		displayLink := strings.Replace(event.ImageLink, "open?", "u/0/uc?", 1)

		imageLine := fmt.Sprintf("![%s](%s)", event.Title, displayLink)
		fmt.Fprintln(f, imageLine)
	}

	if len(event.Summary) > 0 {
		fmt.Fprintln(f, fmt.Sprintf("\n%s", event.Summary))
	}

	fmt.Fprintln(f, breakLine)

	printedDateStr := fmt.Sprintf("Date/Time: **%s.**", event.DateTime.Format(printDateLayout))
	fmt.Fprintln(f, printedDateStr)

	if location := event.location(); len(location) > 0 {
		fmt.Fprintln(f, "")
		printedLocStr := fmt.Sprintf("Location: **%s.**", location)
		fmt.Fprintln(f, printedLocStr)
	}
}

func generateEventPages(events []Event) {
	generateEventList(events)
	generateGroupLog("event")
	for _, event := range events {
		generateEventPage(event.Title, event)
	}
}

func generateExecPage(name string, execs []Exec) {
	generateLog(fmt.Sprintf("%s team", name))

	f, err := os.Create(fmt.Sprintf("./content/team/%s.md", strings.ToLower(name)))
	if err != nil {
		generateErrorLog(fmt.Sprintf("%s team", name))
	}
	defer f.Close()

	generatePageHeader(f, fmt.Sprintf("%s Department", name), "0001-01-01", "", []string{"Team"})
	for _, exec := range execs {
		if exec.Retired >= 0 {
			continue
		}

		var line string

		if exec.PreferredName != "" {
			line = fmt.Sprintf("%s (%s) %s",
				exec.FirstName,
				exec.PreferredName,
				exec.LastName)
		} else {
			line = fmt.Sprintf("%s %s",
				exec.FirstName,
				exec.LastName)
		}

		for i := 0; i < 6; i++ {
			if str := exec.getLink(i); len(str) > 0 {
				line = fmt.Sprintf("[%s](%s)", line, str)
				break
			}
		}

		line = fmt.Sprintf("%s, %s", line, exec.Position)

		if strings.Index(exec.Position, "VP") >= 0 ||
			strings.Index(exec.Position, "President") >= 0 {
			line = "**" + line + "**"
		}

		line = "- " + line

		fmt.Fprintln(f, line)
	}

	if err := f.Close(); err != nil {
		generateErrorLog(fmt.Sprintf("%s team", name))
	}

}

func generateExecPages(execs []Exec) {
	generateGroupLog("exec")
	depts := map[string][]Exec{}
	for _, dept := range getDepartments() {
		depts[dept] = []Exec{}
	}

	for _, exec := range execs {
		for _, dept := range exec.Departments {
			if deptList, exists := depts[dept]; exists {
				depts[dept] = append(deptList, exec)
			}
		}
	}

	for deptName, deptExecs := range depts {
		generateExecPage(deptName, deptExecs)
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

// GenerateEventLinks generates event links for the navbar dropdown meny.
func GenerateEventLinks(events []Event) {
	configFile, err := os.Open(configBase)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	defer configFile.Close()

	lines := []string{}
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i++ {
		if lines[i] != navbar {
			i++
			continue
		}
		i += navbarShift

		preLines := []string{}
		postLines := []string{}
		for j := 0; j < len(lines); j++ {
			if j <= i {
				preLines = append(preLines, lines[j])
			} else {
				postLines = append(postLines, lines[j])
			}
		}

		eventLines := []string{}

		for i := len(events) - 1; i >= 0; i-- {
			filename := events[i].titleToFilename()
			newEvent := []string{
				fmt.Sprintf("        - title: \"%s\"", events[i].Title),
				fmt.Sprintf("          url: /events/%s", filename),
			}

			eventLines = append(newEvent, eventLines...)
		}

		if len(eventLines) > maxNavbarEvents*2 {
			eventLines = eventLines[:maxNavbarEvents*2]
		}

		lines = append(preLines, eventLines...)
		lines = append(lines, postLines...)

		configFile, err := os.Create(config)
		if err != nil {
			log.Fatal(err)
		}
		configWrite := bufio.NewWriter(configFile)
		for _, line := range lines {
			configWrite.WriteString(line + "\n")
		}

		configWrite.Flush()
		configFile.Close()

		break
	}
}

func generateEventList(events []Event) {
	eventsFile, err := os.Create(eventsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range []string{
		"---",
		"title: Events",
		"date: 0001-01-04",
		"sidebar: true",
		"sidebarlogo: whiteside",
		"---",
		"We regularly host events, on our own or in collaboration with other organizations.\n",
		">|Event|>|Date|>|Time|>|Location|",
		">|-----|-|----|-|----|-|--------|",
	} {
		eventsFile.WriteString(line + "\n")
	}

	for i := 0; i < len(events); i++ {
		title := events[i].Title
		filename := events[i].titleToFilename()
		dateStr := events[i].DateTime.Format(printDateLayout)

		listItem := fmt.Sprintf(">|[%s](%s)||%s||%s||%s|",
			title,
			filename,
			dateStr[:len(dateStr)-6],
			dateStr[len(dateStr)-6:],
			events[i].location(),
		)
		eventsFile.WriteString(listItem + "\n")
	}
	eventsFile.Close()
}

const config = "config.yaml"
const configBase = "config_base.yaml"
const navbar = "  navbar:"
const navbarShift = 2
const eventsDirPath = "./content/events/"
const eventsFilePath = "./content/events.md"
const maxNavbarEvents = 3
const breakLine = "---"
