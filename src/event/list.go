package event

import (
	"fmt"

	"gitlab.com/utmist/utmist.gitlab.io/src/helpers"
)

const eventsFilePath = "./content/events.md"
const eventsBasePath = "./assets/events.md"

// Generate the main events list page (events.md).
func generateEventList(events *[]Event, buildings *map[string]Building) {
	// Get header lines of events.md.
	lines := helpers.ReadFileBase(eventsBasePath, len(*events), 10)

	// Add each event to the list.
	for i := 0; i < len(*events); i++ {
		title := (*events)[i].Title
		filename := helpers.StringToFileName((*events)[i].Title)
		dateStr := (*events)[i].DateTime.Format(helpers.PrintDateLayout)

		location, room := (*events)[i].getLocation(buildings)
		head := fmt.Sprintf("##### **[%s](%s)**", title, filename)
		date := fmt.Sprintf("- _Date/Time_: %s", dateStr)
		if len(location) > 0 {
			location = fmt.Sprintf("- _Location_: %s%s", location, func() string {
				if len(room) == 0 {
					return ""
				}
				return fmt.Sprintf(" %s", room)
			}())
		}

		lines = append(lines, []string{head, date, location, helpers.Breakline}...)
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}
