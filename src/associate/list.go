package associate

// MakeEntryList generates a string of associate entries.
func MakeEntryList(
	associates *map[string]Associate,
	entries *[]Entry) []string {

	list := []string{}

	for _, entry := range *entries {
		associate := (*associates)[entry.Email]
		list = append(list, entry.GetListing(&associate, false))
	}

	return list
}
