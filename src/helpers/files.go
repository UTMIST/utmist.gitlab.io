package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const discordBase = "https://discord.gg/"

// StringToFileName formats a given string to a filename.
func StringToFileName(str string) string {
	// We use lowercase page paths.
	filename := strings.ToLower(strings.ToLower(str))

	// Remove illegal characters from filenames.
	strsToRemove := []string{"'", ":", ",", "(", ")", "@", "#"}
	for _, strToRemove := range strsToRemove {
		filename = strings.Replace(filename, strToRemove, "", -1)
	}
	filename = strings.Replace(filename, "&", "and", -1)
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

// SubstituteString stitches new lines into a substitution pattern.
func SubstituteString(lines, newLines []string, substitution string) []string {
	for i := 0; i < len(lines); i++ {
		if lines[i] != substitution {
			continue
		}

		newLines = append(newLines, lines[i+1:]...)
		return append(lines[:i], newLines...)
	}

	return lines
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
func GetJoinLines(description string) []string {
	lines := []string{description}
	InsertDiscordLink(&lines)

	return lines
}
