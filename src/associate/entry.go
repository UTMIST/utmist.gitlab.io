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

const entryOpenTag = "{{< profilePic/profilePicContainer >}}"
const entryCloseTag = "{{< /profilePic/profilePicContainer >}}"
const maxNameChar = 19
const maxPositionChar = 30

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

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
	return abs(e[i].Level) > abs(e[j].Level)
}

// Method Swap() to implement sort.Sort.
func (e EntryList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// IsExecutive returns whether entry is an executive.
func (e *Entry) IsExecutive() bool {
	return e.Level < 0
}

// GetListing returns a listing for this entry.
func (e *Entry) GetListing(associate *Associate, isExec bool) string {

	name := associate.getName()
	position := e.Position
	linkedin := associate.getTargetLink("linkedin")
	github := associate.getTargetLink("github")
	facebook := associate.getTargetLink("facebook")
	twitter := associate.getTargetLink("twitter")
	gitlab := associate.getTargetLink("gitlab")
	personal := associate.getTargetLink("personal")

	if linkedin != "" {
		linkedin = fmt.Sprintf("linkedin=\"%s\"", linkedin)
	}
	if github != "" {
		github = fmt.Sprintf("github=\"%s\"", github)
	}
	if gitlab != "" {
		gitlab = fmt.Sprintf("gitlab=\"%s\"", gitlab)
	}
	if facebook != "" {
		facebook = fmt.Sprintf("facebook=\"%s\"", facebook)
	}
	if twitter != "" {
		twitter = fmt.Sprintf("twitter=\"%s\"", twitter)
	}
	if personal != "" {
		personal = fmt.Sprintf("personal=\"%s\"", personal)
	}

	return fmt.Sprintf("\t{{< profilePic/profilePic  name=\"%s\" position=\"%s\" %s %s %s %s %s %s profile_pic=\"%s\" >}}",
		name,
		position,
		linkedin,
		github,
		gitlab,
		facebook,
		twitter,
		personal,
		associate.ProfilePicturePath)
}

// MakeEntryList generates a string list of associate entries.
func MakeEntryList(
	associates *map[string]Associate,
	entries *[]Entry,
	isDept bool) []string {

	combinedEntries := combineEntries(entries)
	sort.Sort(EntryList(combinedEntries))

	list := []string{}
	list = append(list, entryOpenTag)
	for _, entry := range combinedEntries {
		associate := (*associates)[entry.Email]
		list = append(list, entry.GetListing(&associate, isDept))
	}
	list = append(list, entryCloseTag)
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
	if abs(lvlA) > abs(lvlB) {
		return lvlA
	}
	return lvlB
}
