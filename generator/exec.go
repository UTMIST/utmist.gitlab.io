package generator

import (
	"fmt"
	"os"
	"strings"
)

const websiteLink = 0
const linkedinLink = 1
const gitlabLink = 2
const gitlhubLink = 3
const facebookLink = 4
const twitterLink = 5

const acd = "Academics"
const com = "Communications"
const ext = "External"
const fin = "Finance"
const lgs = "Logistics"
const mkt = "Marketing"
const osg = "Oversight"

// Exec represents an executive member's database row.
type Exec struct {
	Email       string
	PhoneNumber string

	FirstName     string
	PreferredName string
	LastName      string

	Discipline     string
	ProfilePicture string

	Departments []string
	Position    string
	VP          bool

	Website  string
	LinkedIn string
	GitLab   string
	GitHub   string
	Facebook string
	Twitter  string

	Retired int
}

func (e *Exec) getLink(link int) string {
	switch link {
	case websiteLink:
		if len(e.Website) > 0 {
			return e.Website
		}
	case linkedinLink:
		if len(e.LinkedIn) > 0 {
			return fmt.Sprintf("https://linkedin.com/in/%s/", e.LinkedIn)
		}
	case gitlabLink:
		if len(e.GitLab) > 0 {
			return fmt.Sprintf("https://www.gitlab.com/%s/", e.GitLab)
		}
	case gitlhubLink:
		if len(e.GitHub) > 0 {
			return fmt.Sprintf("https://www.github.com/%s/", e.GitHub)
		}
	case facebookLink:
		if len(e.Facebook) > 0 {
			return fmt.Sprintf("https://www.facebook.com/%s/", e.Facebook)
		}
	case twitterLink:
		if len(e.Twitter) > 0 {
			return fmt.Sprintf("https://www.twitter.com/%s/", e.Twitter)
		}
	}

	return ""
}

func getDepartments() []string {
	return []string{
		acd, com, ext, fin, lgs, mkt, osg,
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
