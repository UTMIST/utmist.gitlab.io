package fetcher

import "fmt"

// ASSOCIATES key.
const ASSOCIATES = "ASSOCIATE"

// EVENTS key.
const EVENTS = "EVENT"

// PROJECTS key.
const PROJECTS = "PROJECT"

// RECRUIT key.
const RECRUIT = "RECRUIT"

// Sheet represents the config data for a given sheet.
type Sheet struct {
	ID    string
	Range string
}

// GetSheetKeys returns a sheet ID and range corresponding to sheetname.
func getSheetKeys(sheetName string) (string, string) {
	return fmt.Sprintf("%s_SHEET_ID", sheetName), fmt.Sprintf("%s_SHEET_RANGE", sheetName)
}

// GetSheetNameList returns a list of constant sheet name strings.
func getSheetNameList() []string {
	return []string{
		EVENTS,
		ASSOCIATES,
		PROJECTS,
		RECRUIT,
	}
}
