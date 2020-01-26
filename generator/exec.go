package generator

// Exec represents an executive member's database row.
type Exec struct {
	Email       string
	PhoneNumber string

	FirstName     string
	PreferredName string
	LastName      string

	Discipline string

	Position    string
	VP          bool
	Departments []string

	ProfilePicture   string
	ProfileLink      string
	FacebookUsername string
	TwitterUsername  string
	LinkedInUsername string
	GitHub           string

	Retired int
}

func getDepartments() []string {
	return []string{
		acd, com, ext, fin, lgs, mkt, osg,
	}
}

const acd = "Academics"
const com = "Communications"
const ext = "External"
const fin = "Finance"
const lgs = "Logistics"
const mkt = "Marketing"
const osg = "Oversight"
