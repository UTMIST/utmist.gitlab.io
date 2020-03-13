package associate

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const alm = "Alumni"

// GetDepartmentNames returns a list of department names.
func GetDepartmentNames() []string {
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

	return append(depts, alm)
}
