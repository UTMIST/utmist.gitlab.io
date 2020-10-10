package associate

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const associateRowLength = 14
const entryRowLength = 6
const profilePicPrefix = "https://drive.google.com/uc?export=download&id="
const profilePicPathPrefix = "/images/profilepics/"
const profilePicPathSuffix = ".jpg"
const googleDomain = "google.com"

// LoadAssociate loads an associate from a spreadsheet row.
func LoadAssociate(data []interface{}) Associate {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < associateRowLength; i++ {
		data = append(data, "")
	}
	name := data[0].(string) + data[1].(string) + data[2].(string)
	return Associate{
		data[0].(string),
		data[1].(string),
		data[2].(string),
		data[3].(string),
		data[4].(string),
		data[5].(string),
		data[6].(string),
		processProfilePicLink(data[7].(string), name),
		data[8].(string),
		data[9].(string),
		data[10].(string),
		data[11].(string),
		data[12].(string),
		data[13].(string),
	}
}

// LoadEntries loads an associate entry from a spreadsheet row.
func LoadEntries(data []interface{}, associates *map[string]Associate) []Entry {

	// Pad the columns with blanks to avoid index-out-of-range.
	for i := len(data); i < entryRowLength; i++ {
		data = append(data, "")
	}
	positions := strings.Split(data[4].(string), ",")
	departments := strings.Split(data[5].(string), ",")

	entries := []Entry{}

	levelRegexes := helpers.GetPosRanks()
	execRegexes := helpers.GetPosExec()

	for i := 0; i < len(positions); i++ {
		if i >= len(departments) {
			break
		}

		email := data[3].(string)
		associate := (*associates)[email]

		level := 0 //associate level (lowest)
		posTrimmed := strings.TrimSpace(positions[i])

		for r, regStr := range levelRegexes { //determine the level
			if helpers.FitRegex(posTrimmed, regStr) {
				level = r + 1
				break
			}
		}

		if level != 0 {
			for _, regStr := range execRegexes { //check if they are exec
				if helpers.FitRegex(posTrimmed, regStr) {
					level = 0 - level
					break
				}
			}
		}

		entries = append(
			entries,
			Entry{
				email,
				posTrimmed,
				strings.TrimSpace(departments[i]),
				&associate,
				level})
	}

	return entries
}

func processProfilePicLink(link string, name string) string {
	if strings.Contains(link, googleDomain) {
		return ""
	}
	return downloadProfilePic(link, name)
}

func downloadProfilePic(link string, name string) string {
	response, err := http.Get(link)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	downloadPath := fmt.Sprintf("static%s%s%s",
		profilePicPathPrefix,
		name,
		profilePicPathSuffix)
	img, err := os.Create(downloadPath)
	if err != nil {
		return ""
	}
	defer img.Close()
	_, err = io.Copy(img, response.Body)
	if err != nil {
		return ""
	}
	filepath := fmt.Sprintf("%s%s%s",
		profilePicPathPrefix,
		name,
		profilePicPathSuffix)

	return filepath
}
