package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const discordBase = "https://discord.gg/"
const joinParagraphFile = "assets/join.md"

// StringToFileName formats a given string to a filename.
func StringToFileName(str string) string {
	// We use lowercase page paths.
	filename := strings.ToLower(strings.ToLower(str))

	// Remove illegal characters from filenames.
	strsToRemove := []string{"'", ":", ",", "(", ")", "@"}
	for _, strToRemove := range strsToRemove {
		filename = strings.Replace(filename, strToRemove, "", -1)
	}
	filename = strings.Replace(filename, " - ", " ", -1)
	filename = strings.Replace(filename, "  ", " ", -1)
	filename = strings.Replace(filename, " ", "-", -1)

	return filename
}

// ReadContentLines reads lines from a content file.
func ReadContentLines(filename string) []string {
	eventsFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer eventsFile.Close()

	// Read lines from config_base.
	lines := []string{}
	scanner := bufio.NewScanner(eventsFile)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// OverwriteWithLines overwrites the given file at <path> with <lines>.
func OverwriteWithLines(path string, lines []string) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		file.WriteString(line + "\n")
	}
	file.Close()
}

// StitchIntoLines stitches new lines into the config.
func StitchIntoLines(lines, newLines *[]string, start string, shift int) {

	// Search for correct place
	startIndex := 0
	for startIndex < len(*lines) {
		if (*lines)[startIndex] == start {
			break
		}
		startIndex++
	}
	startIndex += shift

	// Store the lines that go before/after.
	preLines := []string{}
	for j := 0; j <= startIndex; j++ {
		preLines = append(preLines, (*lines)[j])
	}
	postLines := []string{}
	for j := startIndex + 1; j < len(*lines); j++ {
		postLines = append(postLines, (*lines)[j])
	}

	// Stitch config.yaml back together with preLines and postLines.
	*lines = append(preLines, (*newLines)...)
	*lines = append(*lines, postLines...)
}

// ReadFileBase an existing file and truncates it to the header.
func ReadFileBase(filename string, num, trunc int) []string {

	lines := ReadContentLines(filename)
	for i, line := range lines {
		lines[i] = line
	}

	if trunc == -1 {
		return lines
	}

	return lines[:trunc]
}

// InsertDiscordLink appends the server invite link to Discord links.
func InsertDiscordLink(lines *[]string) {
	discordLink, exists := os.LookupEnv("DISCORD_LINK")
	if !exists {
		return
	}

	for i := range *lines {
		(*lines)[i] = strings.Replace((*lines)[i],
			discordBase,
			fmt.Sprintf("%s%s", discordBase, discordLink),
			-1)
	}
}

// GetJoinLines returns the lines from the join prompt paragraph.
func GetJoinLines() []string {
	lines := ReadContentLines(joinParagraphFile)
	InsertDiscordLink(&lines)

	return lines
}
