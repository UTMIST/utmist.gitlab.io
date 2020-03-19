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
	lines := helpers.ReadFileBase(eventsBasePath, len(*events), 11)

	// Add each event to the list.
	for i := 0; i < len(*events); i++ {
		title := (*events)[i].Title
		filename := helpers.StringToFileName((*events)[i].Title)
		dateStr := (*events)[i].DateTime.Format(helpers.PrintDateLayout)

		location, room := (*events)[i].getLocation(buildings)

		eventListing := fmt.Sprintf("|[%s](%s)|%s|%s|%s|%s|\n|%s|%s|%s|%s|%s|",
			title,
			filename,
			helpers.TablePadder,
			dateStr[:len(dateStr)-6],
			helpers.TablePadder,
			location,
			"",
			helpers.TablePadder,
			dateStr[len(dateStr)-6:],
			helpers.TablePadder,
			room,
		)
		lines = append(lines, eventListing)
		lines = append(lines, fmt.Sprintf("||%s||%s||", helpers.TablePadder, helpers.TablePadder))
	}

	helpers.OverwriteWithLines(eventsFilePath, lines)
}
