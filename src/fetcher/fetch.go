package fetcher

import (
	"fmt"
	"log"
	"os"
	"sort"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// SCOPE of the sheets API access.
const SCOPE = "https://www.googleapis.com/auth/spreadsheets.readonly"

// Fetch fetches associate, event, project, recruitment databases.
func Fetch() (
	[]event.Event,
	[]associate.Associate,
	[]position.Position,
	[]project.Project,
	[]project.Project) {

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
	positions := []position.Position{}
	activeProjects := []project.Project{}
	pastProjects := []project.Project{}

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
		case POSITIONS:
			for _, row := range resp.Values {
				positions = append(positions, position.LoadPosition(row))
			}
		case PROJECTS:
			for _, row := range resp.Values {
				proj := project.LoadProject(row)
				if proj.Status == project.ActiveStatus {
					activeProjects = append(activeProjects, proj)
				} else {
					pastProjects = append(pastProjects, proj)
				}
			}
		default:
			fmt.Printf("Fetch for %s not yet implemented.\n", sheetName)
		}
	}

	sort.Sort(event.List(events))

	return events, associates, positions, activeProjects, pastProjects
}

// LoadFetchEnv loads environment variables from .env for fetching.
func loadFetchEnv() (map[string]Sheet, error) {
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
