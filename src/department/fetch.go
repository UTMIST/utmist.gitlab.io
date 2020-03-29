package department

import (
	"fmt"
	"os"
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
)

const departmentsSheetRange = 2

// GetDeptNames returns a list of department names.
func GetDeptNames(alumni bool) []string {
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
		depts = append(depts, associate.ALM)
	}
	return depts
}

// LoadDeptDescs returns a map of department descriptions.
func LoadDeptDescs(deptDescs *map[string]string, data []interface{}) {
	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < departmentsSheetRange; i++ {
		data = append(data, "")
	}

	// Map the department to its description.
	dept := data[0].(string)
	desc := data[1].(string)

	(*deptDescs)[dept] = desc
}
