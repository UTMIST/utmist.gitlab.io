package associate

import (
	"fmt"
	"sort"
	"strings"
)

// Entry represents one position listing for an associate.
// 0 - lowest level, normal associate; natural number: bolded (dept.) non-exec; negative integer: bolded exec
type Entry struct {
	Email      string
	Position   string
	Department string
	Associate  *Associate
	Level      int
}

// EntryList defines a list of entries.
type EntryList []Entry

// level value reserved for roles bolded on the exec list (currently only the president)
const topLevel = -1

// Method Len() to implement sort.Sort.
func (e EntryList) Len() int {
	return len(e)
}

// Method Less() to implement sort.Sort.
func (e EntryList) Less(i, j int) bool {
	// if the levels are equal, sort alphabetically
	if e[i].Level == e[j].Level {
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
	}

	// 0 is the lowest rank, not the highest
	if e[i].Level != 0 && e[j].Level == 0 {
		return true
	}
	if e[i].Level == 0 && e[j].Level != 0 {
		return false
	}

	// compare the absolute values of the levels
	lvlI := e[i].Level
	if lvlI < 0 {
		lvlI = -lvlI
	}

	lvlJ := e[j].Level
	if lvlJ < 0 {
		lvlJ = -lvlJ
	}

	// lower levels means higher rank
	if lvlI < lvlJ {
		return true
	}
	return false
}

// Method Swap() to implement sort.Sort.
func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// IsExecutive returns whether entry is an executive.
func (e *Entry) IsExecutive() bool {
	return e.Level < 0
}

// IsToBeBolded returns whether listing should be bolded as Executive.
func (e *Entry) IsToBeBolded(dept bool) bool {
	// Bold if specified in .env POS_RANKING, or if top position (level == topLevel).
	return (dept && e.isSignificant()) || e.isTop()
}

// isSignificant returns whether this position would be bolded on its department page
func (e *Entry) isSignificant() bool {
	return e.Level != 0
}

func (e *Entry) isTop() bool {
	return e.Level == topLevel
}

func (e *Entry) isVP() bool {
	return strings.HasPrefix(e.Position, "VP")
}

func (e *Entry) isAVP() bool {
	return strings.HasPrefix(e.Position, "AVP")
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
			newEntry.Level = combineLevel(existingEntry.Level, e.Level)
		}

		entryMap[e.Email] = newEntry
	}

	combinedEntries := []Entry{}
	for _, entry := range entryMap {
		combinedEntries = append(combinedEntries, entry)
	}

	return combinedEntries
}

//combineLevel determines the highest level for an associate
func combineLevel(lvlA, lvlB int) int {
	// if the levels are the same
	if lvlA == lvlB {
		return lvlA
	}
	// if one of the levels is the lowest (associate 0)
	if lvlA == 0 {
		return lvlB
	}
	if lvlB == 0 {
		return lvlA
	}

	isExec := lvlA < 0 || lvlB < 0

	// exec levels are negative
	if lvlA < 0 {
		lvlA = -lvlA
	}
	if lvlB < 0 {
		lvlB = -lvlB
	}
	combinedLevel := lvlA

	// lower levels are ranked higher
	if lvlA > lvlB {
		combinedLevel = lvlB
	}

	if isExec {
		combinedLevel = -combinedLevel
	}
	return combinedLevel
}
