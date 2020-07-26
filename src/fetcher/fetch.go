package fetcher

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// SCOPE of the sheets API access.
const SCOPE = "https://www.googleapis.com/auth/spreadsheets.readonly"

// Fetch fetches associate, event, project, recruitment databases.
func Fetch() (map[string]associate.Associate, map[int][]associate.Entry) {

	b, err := getCredentials()
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	b = append(b, 10)

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, SCOPE)
	if err != nil {
		log.Print(string(b))
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	// Create service.
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	firstYear, lastYear := helpers.GetYearRange()

	associates := fetchAssociates(srv)
	entries := fetchEntries(srv, firstYear, lastYear)

	return associates, entries
}

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

func fetchEntries(srv *sheets.Service, firstYear, lastYear int) map[int][]associate.Entry {

	entries := map[int][]associate.Entry{}
	sheetID := os.Getenv("ASSOCIATES_SHEET_ID")
	for year := firstYear; year <= lastYear; year++ {
		yearEntries := []associate.Entry{}
		sheetRange := os.Getenv(fmt.Sprintf("ASSOCIATES_%d", year))
		resp := fetchValues(
			srv,
			fmt.Sprintf("Associates (%d)", year),
			sheetID,
			sheetRange)
		for _, row := range resp.Values {
			yearEntries = append(yearEntries, associate.LoadEntries(row)...)
		}

		entries[year] = yearEntries
	}

	return entries
}

func fetchValues(srv *sheets.Service, groupName, sheetID, sheetRange string) *sheets.ValueRange {
	// Validate the API response.
	resp, err := srv.Spreadsheets.Values.Get(sheetID, sheetRange).Do()
	if err != nil {
		log.Println(fmt.Sprintf("Unable to retrieve %s data from sheet: ",
			groupName))
		panic(err)
	}
	if len(resp.Values) == 0 {
		log.Printf("No %s data found.\n", groupName)
	}

	log.Printf("Downloaded %s data.", groupName)

	return resp
}
