package associate

import "sort"

// MakeEntryList generates a string of associate entries.
func MakeEntryList(
	associates *map[string]Associate,
	entries *[]Entry,
	isDept bool) []string {

	sort.Sort(EntryList(*entries))

	list := []string{}
	for _, entry := range *entries {
		associate := (*associates)[entry.Email]
		list = append(list, entry.GetListing(&associate, isDept))
	}

	return list
}
