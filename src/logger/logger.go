package logger

import (
	"fmt"
	"log"
)

func GenerateLog(str string) {
	log.Println(fmt.Sprintf("\tGenerating page for %s.", str))
}

func GenerateErrorLog(str string) {
	log.Println(fmt.Sprintf("\tFailed to generate page for %s.", str))
}

func GenerateGroupLog(str string) {
	log.Println(fmt.Sprintf("Generating %s pages.", str))
}
