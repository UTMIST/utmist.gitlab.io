package event

import "gitlab.com/utmist/utmist.gitlab.io/src/helpers"

// Where the event list starts in the config.
const start = "    - title: Events"

// Dictating how many individual links appear on the navbar list.
const max = 3

// GenerateNavbarEventLinks generates event links for the navbar dropdown menu.
func GenerateNavbarEventLinks(events []Event, lines *[]string) {
	eventTitles := []string{}
	for i := max - 1; i > -1; i-- {
		eventTitles = append([]string{events[i].Title}, eventTitles...)
	}

	helpers.StitchPageLink(lines, eventTitles, "/events/", start)
}
