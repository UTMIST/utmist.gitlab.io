package associate

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const alm = "Alumni"

// Interface to allow groupByDept to operate on structs by calling .Dept()
type hasDepartment interface {
	Dept() string
}

// GetDepartmentNames returns a list of department names.
func GetDepartmentNames(alumni bool) []string {
	year, exists := os.LookupEnv(("DEPTS_YEAR"))
	if !exists {
		year = fmt.Sprintf("%d", time.Now().Year())
	}

	depts := []string{}
	envDepts, exists := os.LookupEnv(fmt.Sprintf("DEPTS_%s", year))
	if exists {
		for _, d := range strings.Split(envDepts, ",") {
			depts = append(depts, d)
		}
	}
	if alumni {
		depts = append(depts, alm)
	}
	return depts
}

// Dept implements hasDepartment().
func (a Associate) Dept() string {
	return a.Department
}
