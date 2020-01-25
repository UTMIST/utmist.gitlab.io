package generator

// Exec represents an executive member's database row.
type Exec struct {
	Email       string
	PhoneNumber int

	FirstName     string
	PreferredName string
	LastName      string

	Discipline string

	Position    string
	VP          bool
	Departments string

	ProfilePicture   string
	ProfileLink      string
	FacebookUsername string
	TwitterUsername  string
	LinkedInUsername string
	GitHub           string

	Retired int
}
