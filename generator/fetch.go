package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// SCOPE of the sheets API access.
const SCOPE = "https://www.googleapis.com/auth/spreadsheets.readonly"

// Fetch fetches events, projects, recruitment, and exec databases.
func Fetch() ([]Event, []Exec, []Project) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, SCOPE)
	if err != nil {
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
	events := []Event{}
	execs := []Exec{}
	projects := []Project{}

	for _, sheetName := range getSheetNameList() {
		sheet, exists := sheets[sheetName]
		if !exists {
			continue
		}

		resp, err := srv.Spreadsheets.Values.Get(sheet.ID, sheet.Range).Do()
		if err != nil {
			dataErrorLog(sheetName, err)
			continue
		} else if len(resp.Values) == 0 {
			dataNoneLog(sheetName)
			continue
		} else {
			dataSuccessLog(sheetName)
		}

		switch sheetName {
		default:
			fetchLog(fmt.Sprintf("Fetch for %s not yet implemented.", sheetName))
		}
	}

	return events, execs, projects
}

func loadEvents()   {}
func loadExecs()    {}
func loadProjects() {}

// LoadFetchEnv loads environment variables from .env for fetching.
func loadFetchEnv() (map[string]Sheet, error) {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return nil, err
	}

	sheetIDRanges := map[string]Sheet{}
	sheetNames := getSheetNameList()

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
