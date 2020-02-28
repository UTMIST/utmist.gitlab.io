package fetcher

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/associate"
	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/project"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

const eventSheetRange = 8
const associateSheetRange = 17
const parseDateLayout = "01/02/2006 15:04:05"

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
			dataErrorLog(sheetName, err)
			continue
		}
		if len(resp.Values) == 0 {
			dataNoneLog(sheetName)
			continue
		}

		dataSuccessLog(sheetName)

		// Add to the appropriate list.
		switch sheetName {
		case ASSOCIATES:
			for _, row := range resp.Values {
				associates = append(associates, loadAssociate(row))
			}
		case EVENTS:
			for _, row := range resp.Values {
				events = append(events, loadEvent(row))
			}
		default:
			fetchLog(fmt.Sprintf("Fetch for %s not yet implemented.", sheetName))
		}
	}

	return events, associates, projects
}

// Load an event from a spreadsheet row.
func loadEvent(data []interface{}) event.Event {
	for i := len(data); i < eventSheetRange; i++ {
		data = append(data, "")
	}

	event := event.Event{
		Title:     data[0].(string),
		Type:      data[1].(string),
		DateTime:  formatDateEST(data[2].(string)),
		Location:  data[3].(string),
		Summary:   data[4].(string),
		ImageLink: data[5].(string),
		PreLink:   data[6].(string),
		PostLink:  data[7].(string),
	}

	return event
}

// Load an associate from a spreadsheet row.
func loadAssociate(data []interface{}) associate.Associate {
	for i := len(data); i < associateSheetRange; i++ {
		data = append(data, "")
	}

	associate := associate.Associate{
		Email:         data[0].(string),
		FirstName:     data[1].(string),
		PreferredName: data[2].(string),
		LastName:      data[3].(string),
		PhoneNumber:   data[4].(string),
		Position:      data[5].(string),
		Departments:   strings.Split(data[6].(string), ", "),
		Discipline:    data[8].(string),
		Website:       data[10].(string),
		Facebook:      data[11].(string),
		Twitter:       data[12].(string),
		LinkedIn:      data[13].(string),
		GitHub:        data[14].(string),
		GitLab:        data[15].(string),
		Retired: func(yearObj interface{}) int {
			retired, err := strconv.Atoi(yearObj.(string))
			if err != nil {
				retired = -1
			}
			return retired
		}(data[16]),
	}
	return associate
}
func loadProject() {}

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

// Need to save logs. TODO
func dataErrorLog(str string, err error) {
	fetchLogFatal(fmt.Sprintf("Unable to retrieve %s data from sheet.", str), err)
}

func dataSuccessLog(str string) {
	fetchLog(fmt.Sprintf("Downloaded %s data.", str))
}

func dataNoneLog(str string) {
	fetchLog(fmt.Sprintf("No %s data found.", str))
}

func fetchLogFatal(str string, err error) {
	log.Println(str)
}

func fetchLog(str string) {
	log.Println(str)
}

// Format a date from EST.
func formatDateEST(dateStr string) time.Time {
	// UNDOCUMENTED.
	parts := strings.Split(dateStr, "/")
	for i := 0; i < 2; i++ {
		if len(parts[i]) == 1 {
			parts[i] = fmt.Sprintf("0%s", parts[i])
		}
	}
	dateStr = strings.Join(parts, "/")

	// UNDOCUMENTED.
	parts = strings.Split(dateStr, " ")
	if parts[1][1] == ':' {
		parts[1] = fmt.Sprintf("0%s", parts[1])
	}
	dateStr = strings.Join(parts, " ")

	// Load location and parse the time in that location.
	toronto, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}
	}
	dateTime, err := time.ParseInLocation(parseDateLayout, dateStr, toronto)
	if err != nil {
		return time.Time{}
	}

	return dateTime
}
