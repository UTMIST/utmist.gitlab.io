package associate

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/hugo"
	"gitlab.com/utmist/utmist.gitlab.io/src/logger"
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

// Associate represents an associateutive member's database row.
type Associate struct {
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

// Return selected social media link for user.
func (e *Associate) getLink(link int) string {
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

// Get list of departments.
func getDepartments() []string {
	return []string{
		acd, com, ext, fin, lgs, mkt, osg,
	}
}

// Generate a page for the a department.
func generateDepartmentPage(name string, associates []Associate) {
	logger.GenerateLog(fmt.Sprintf("%s team", name))

	// Create file for the page and write the header.
	f, err := os.Create(fmt.Sprintf("./content/team/%s.md", strings.ToLower(name)))
	if err != nil {
		logger.GenerateErrorLog(fmt.Sprintf("%s team", name))
	}
	defer f.Close()
	hugo.GeneratePageHeader(f, fmt.Sprintf("%s Department", name), "0001-01-01", "", []string{"Team"})

	// Write a list entry for every member; skip the alumni (retired).
	for _, associate := range associates {
		if associate.Retired >= 0 {
			continue
		}

		var line string
		if associate.PreferredName != "" {
			line = fmt.Sprintf("%s (%s) %s",
				associate.FirstName,
				associate.PreferredName,
				associate.LastName)
		} else {
			line = fmt.Sprintf("%s %s",
				associate.FirstName,
				associate.LastName)
		}

		// Pick the first available social media link from the defined order.
		for i := 0; i < 6; i++ {
			if str := associate.getLink(i); len(str) > 0 {
				line = fmt.Sprintf("[%s](%s)", line, str)
				break
			}
		}

		// Reformat the line and write it.
		line = fmt.Sprintf("%s, %s", line, associate.Position)
		if strings.Index(associate.Position, "VP") >= 0 ||
			strings.Index(associate.Position, "President") >= 0 {
			line = "**" + line + "**"
		}
		line = "- " + line
		fmt.Fprintln(f, line)
	}

	// Try closing the file.
	if err := f.Close(); err != nil {
		logger.GenerateErrorLog(fmt.Sprintf("%s team", name))
	}

}

// GenerateAssociatePages generates all the department pages.
func GenerateAssociatePages(associates []Associate) {
	logger.GenerateGroupLog("associate")

	// Populate the departments with empty associate list.
	depts := map[string][]Associate{}
	for _, dept := range getDepartments() {
		depts[dept] = []Associate{}
	}

	// Load associates into their department's associate list.
	for _, associate := range associates {
		for _, dept := range associate.Departments {
			if deptList, exists := depts[dept]; exists {
				depts[dept] = append(deptList, associate)
			}
		}
	}

	// Generate each department page.
	for deptName, deptAssociates := range depts {
		generateDepartmentPage(deptName, deptAssociates)
	}
}
