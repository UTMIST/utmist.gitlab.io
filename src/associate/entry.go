package associate

import (
	"fmt"
	"sort"
	"strings"
)

// Entry represents one position listing for an associate.
type Entry struct {
	Email      string
	Position   string
	Department string
	Associate  *Associate
}

// EntryList defines a list of entries.
type EntryList []Entry

// Method Len() to implement sort.Sort.
func (e EntryList) Len() int {
	return len(e)
}

// Method Less() to implement sort.Sort.
func (e EntryList) Less(i, j int) bool {
	if !e[i].IsExecutive() && e[j].IsExecutive() {
		return false
	}
	for _, criteria := range []int{
		strings.Compare(e[i].Position, e[j].Position),
		strings.Compare(e[i].Associate.Surname, e[j].Associate.Surname),
		strings.Compare(e[i].Associate.GivenName, e[j].Associate.GivenName),
		strings.Compare(e[i].Associate.PreferredName, e[j].Associate.PreferredName)} {
		switch criteria {
		case -1:
			return true
		case 1:
			return false
		}
	}

	return false
}

// Method Swap() to implement sort.Sort.
func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// IsExecutive returns whether entry is an executive.
func (e *Entry) IsExecutive() bool {
	return e.isVP() || e.isPresident()
}

// IsToBeBolded returns whether listing should be bolded as Executive.
func (e *Entry) IsToBeBolded(dept bool) bool {
	// Bold if VP on department page, or if (Co-)President.
	return (dept && e.isVP()) || e.isPresident()
}

func (e *Entry) isPresident() bool {
	return strings.Index(e.Position, "President") >= 0
}

func (e *Entry) isVP() bool {
	return strings.Index(e.Position, "VP") >= 0
}

// GetListing returns a listing for this entry.
func (e *Entry) GetListing(associate *Associate, isExec bool) string {

	var listing string

	name := associate.getName()
	if link := associate.getLink(); len(link) > 0 {
		listing = fmt.Sprintf("[%s](%s), %s", name, link, e.Position)
	} else {
		listing = fmt.Sprintf("%s, %s", name, e.Position)
	}

	if e.IsToBeBolded(isExec) {
		return fmt.Sprintf("- **%s**", listing)
	}
	return fmt.Sprintf("- %s", listing)
}

// MakeEntryList generates a string list of associate entries.
func MakeEntryList(
	associates *map[string]Associate,
	entries *[]Entry,
	isDept bool) []string {

	combinedEntries := combineEntries(entries)
	sort.Sort(EntryList(combinedEntries))

	list := []string{}
	for _, entry := range combinedEntries {
		associate := (*associates)[entry.Email]
		list = append(list, entry.GetListing(&associate, isDept))
	}

	return list
}

func combineEntries(entries *[]Entry) []Entry {

	entryMap := map[string]Entry{}
	for _, e := range *entries {

		newEntry := e
		if existingEntry, exists := entryMap[e.Email]; exists {
			newEntry.Position = fmt.Sprintf("%s, %s", existingEntry.Position, e.Position)
		}

		entryMap[e.Email] = newEntry
	}

	combinedEntries := []Entry{}
	for _, entry := range entryMap {
		combinedEntries = append(combinedEntries, entry)
	}

	return combinedEntries
}
