package fetcher

import (
	"fmt"
	"log"
	"os"
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// SCOPE of the sheets API access.
const SCOPE = "https://www.googleapis.com/auth/spreadsheets.readonly"

// Fetch fetches associate, event, project, recruitment databases.
func Fetch() ([]event.Event, []associate.Associate, []project.Project) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

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

	// Pull sheet IDs and ranges.
	sheets, err := loadFetchEnv()
	if err != nil {
		log.Fatalf("Unable to load sheet IDs: %v", err)
	}

	// Create lists.
	events := []event.Event{}
	associates := []associate.Associate{}
	projects := []project.Project{}

	// Populate each list with associates, events, project, respectively.
	for _, sheetName := range getSheetNameList() {
		sheet, exists := sheets[sheetName]
		if !exists {
			continue
		}

		// Validate the API response.
		resp, err := srv.Spreadsheets.Values.Get(sheet.ID, sheet.Range).Do()
		if err != nil {
			log.Println(fmt.Sprintf("Unable to retrieve %s data from sheet: ", sheetName))
			continue
		}
		if len(resp.Values) == 0 {
			log.Printf("No %s data found.\n", sheetName)
			continue
		}

		log.Printf("Downloaded %s data.", sheetName)

		// Add to the appropriate list.
		switch sheetName {
		case ASSOCIATES:
			for _, row := range resp.Values {
				associates = append(associates, associate.LoadAssociate(row)...)
			}
		case EVENTS:
			for _, row := range resp.Values {
				events = append(events, event.LoadEvent(row))
			}
		default:
			fmt.Printf("Fetch for %s not yet implemented.\n", sheetName)
		}
	}

	sort.Sort(event.List(events))

	return events, associates, projects
}

// LoadFetchEnv loads environment variables from .env for fetching.
func loadFetchEnv() (map[string]Sheet, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	sheetIDRanges := map[string]Sheet{}
	sheetNames := getSheetNameList()

	// Populates sheet IDs and ranges for each group (associates, vents, projects)
	for _, sheetName := range sheetNames {
		ID, Range := getSheetKeys(sheetName)
		sheetID, ok := os.LookupEnv(ID)
		if !ok {
			continue
		}
		sheetRange, ok := os.LookupEnv(Range)
		if !ok {
			continue
		}
		sheetIDRanges[sheetName] = Sheet{
			ID:    sheetID,
			Range: sheetRange,
		}
	}

	return sheetIDRanges, nil
}
