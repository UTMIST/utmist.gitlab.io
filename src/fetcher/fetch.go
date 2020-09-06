package fetcher

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// SCOPE of the sheets API access
const SCOPE = "https://www.googleapis.com/auth/spreadsheets.readonly"

// FetchFromGoogleSheets pulls associate/recruitment data from Google Sheets.
func FetchFromGoogleSheets() (
	map[string]associate.Associate,
	map[int][]associate.Entry,
	[]position.Position) {

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

	firstYear, lastYear, err := helpers.GetYearRange(os.Getenv("YEARS"))
	if err != nil {
		panic(err)
	}

	associates := fetchAssociates(srv)
	entries := fetchAssociateEntries(srv, &associates, firstYear, lastYear)
	positions := fetchPositions(srv)

	return associates, entries, positions
}

// FetchFromOneDriveFiles pulls event/project data from local OneDrive files.
func FetchFromOneDriveFiles() (
	map[int][]event.Event,
	map[int][]project.Project) {

	events := fetchEvents()
	projects := fetchProjects()

	return events, projects
}

func fetchValues(
	srv *sheets.Service,
	groupName,
	sheetID,
	sheetRange string) *sheets.ValueRange {

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
