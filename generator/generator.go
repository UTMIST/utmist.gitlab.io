package generator

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func generateEventPage() {

}

func generateEventPages() {

}

func generateExecPage(name string, execs []Exec) {
	generateExecLog(name)
	f, err := os.Create(fmt.Sprintf("./content/team/%s.md", strings.ToLower(name)))
	if err != nil {
		generateExecLogError(name)
	}
	header := []string{
		"---",
		fmt.Sprintf("title: %s Department", name),
		"date: 0001-01-01",
		"sidebar: true",
		"sidebarlogo: whiteside",
		"---",
		"",
	}

	for _, line := range header {
		fmt.Fprintln(f, line)
	}

	for _, exec := range execs {
		var line string

		if exec.PreferredName != "" {
			line = fmt.Sprintf("%s (%s) %s, %s",
				exec.FirstName,
				exec.PreferredName,
				exec.LastName,
				exec.Position)
		} else {
			line = fmt.Sprintf("%s %s, %s",
				exec.FirstName,
				exec.LastName,
				exec.Position)
		}

		if strings.Index(exec.Position, "VP") >= 0 ||
			strings.Index(exec.Position, "President") >= 0 {
			line = "**" + line + "**"
		}

		line = "- " + line

		fmt.Fprintln(f, line)
	}

	if err := f.Close(); err != nil {
		generateExecLogError(name)
	}

}

func generateExecPages(execs []Exec) {
	log.Println("Generating exec pages.")
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
}

func generateExecLog(str string) {
	log.Println(fmt.Sprintf("\tGenerating page for %s team.", str))
}

func generateExecLogError(str string) {
	log.Println(fmt.Sprintf("\tFailed to generate page for %s team.", str))
}
