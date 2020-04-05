package helpers

import (
	"fmt"
	"log"
)

// GenerateLog logs when creating a page.
func GenerateLog(str string) {
	log.Println(fmt.Sprintf("\tGenerating page for %s.", str))
}

// GenerateErrorLog logs when creating a page produces an error.
func GenerateErrorLog(str string) {
	log.Println(fmt.Sprintf("\tFailed to generate page for %s.", str))
}

// GenerateGroupLog logs when generating a group of pages.
func GenerateGroupLog(str string) {
	log.Println(fmt.Sprintf("Generating %s pages.", str))
}
