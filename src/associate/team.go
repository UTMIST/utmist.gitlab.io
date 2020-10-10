package associate

import (
	"strings"
)

// AddTeamEntries addes new entries
func AddTeamEntries(
	addEntries []Entry,
	entries map[string][]Entry) map[string][]Entry {
	for _, entry := range addEntries {

		titles := strings.Split(entry.Position, ",")
		for _, title := range titles {
			if strings.Contains(title, "(") && strings.Contains(title, ")") {
				start := strings.Index(title, "(")
				end := strings.Index(title, ")")
				team := title[start+1 : end]
				position := strings.TrimSpace(title[:start])
				entry.Position = position
				entry.Department = team
				if _, ok := entries[team]; ok {
					entries[team] = append(entries[team], entry)
				} else {
					entries[team] = []Entry{entry}
				}
			}
		}
	}
	return entries
}
