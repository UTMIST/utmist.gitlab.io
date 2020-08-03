package fetcher

import (
	"io/ioutil"
	"log"

	"gitlab.com/utmist/utmist.gitlab.io/src/event"
	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

func fetchEvents() map[int][]event.Event {
	files, err := ioutil.ReadDir(event.EventsFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	events := map[int][]event.Event{}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		if _, _, err := helpers.GetYearRange(f.Name()); err == nil {
			continue
		}

		event := event.LoadEvent(f.Name())
		year := event.DateTime.Year()
		if event.DateTime.Month() < 9 {
			year--
		}

		events[year] = append(events[year], event)
	}

	return events
}
