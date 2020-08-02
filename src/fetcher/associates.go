package fetcher

import (
	"fmt"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"google.golang.org/api/sheets/v4"
)

func fetchAssociates(srv *sheets.Service) map[string]associate.Associate {
	associates := map[string]associate.Associate{}

	sheetID := os.Getenv("ASSOCIATES_SHEET_ID")
	sheetRange := os.Getenv("ASSOCIATES_RANGE")
	resp := fetchValues(srv, "Associates Directory", sheetID, sheetRange)
	for _, row := range resp.Values {
		associate := associate.LoadAssociate(row)
		associates[associate.UofTEmail] = associate
	}

	return associates
}

func fetchAssociateEntries(
	srv *sheets.Service,
	associates *map[string]associate.Associate,
	firstYear,
	lastYear int) map[int][]associate.Entry {

	entries := map[int][]associate.Entry{}
	sheetID := os.Getenv("ASSOCIATES_SHEET_ID")
	for y := firstYear; y <= lastYear; y++ {
		yearEntries := []associate.Entry{}
		sheetRange := fmt.Sprintf("%d!%s", y, os.Getenv("ENTRIES_RANGE"))
		resp := fetchValues(
			srv,
			fmt.Sprintf("Associates (%d)", y),
			sheetID,
			sheetRange)
		for _, row := range resp.Values {
			yearEntries = append(
				yearEntries,
				associate.LoadEntries(row, associates)...)
		}

		entries[y] = yearEntries
	}

	return entries
}
