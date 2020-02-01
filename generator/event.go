package generator

import (
	"fmt"
	"strings"
	"time"
)

// Event represents an entry in the Events Google Sheet
type Event struct {
	Title    string
	Type     string
	DateTime time.Time
	Location string

	Summary   string
	ImageLink string
	PreLink   string
	PostLink  string
}

func (e *Event) titleToFilename() string {
	filename := strings.Replace(strings.ToLower(e.Title), "'", "", -1)
	filename = strings.Replace(filename, ":", "", -1)
	filename = strings.Replace(filename, ",", "", -1)
	filename = strings.Replace(filename, "(", "", -1)
	filename = strings.Replace(filename, ")", "", -1)
	filename = strings.Replace(filename, " - ", " ", -1)
	filename = strings.Replace(filename, " ", "-", -1)

	return filename
}

func (e *Event) location() string {
	campusCodeMap := map[string]string{
		"BA": "080",
		"ES": "062",
		"GB": "070",
		"SF": "009",
		"SS": "033",
	}

	campusNameMap := map[string]string{
		"BA": "Bahen Centre for IT",
		"ES": "Earth Sciences Centre",
		"GB": "Galbraith Building",
		"SF": "Sanford Fleming Building",
		"SS": "Sidney Smith Hall",
	}

	if len(e.Location) > 2 {

		buildingCode := e.Location[:2]

		if number, exists := campusCodeMap[buildingCode]; exists &&
			buildingCode == strings.ToUpper(buildingCode) {

			return fmt.Sprintf("[%s](http://map.utoronto.ca/utsg/building/%s) %s",
				campusNameMap[buildingCode], number, e.Location)
		}
	}

	return e.Location

}
