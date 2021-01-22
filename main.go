package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"gitlab.com/utmist/utmist.gitlab.io/src/points"

	"github.com/joho/godotenv"
	"gitlab.com/utmist/utmist.gitlab.io/src/bundle"
	"gitlab.com/utmist/utmist.gitlab.io/src/fetcher"
	"gitlab.com/utmist/utmist.gitlab.io/src/generator"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	flag.Parse()

	associates, assocEntries, teamEntries, positions, tasks, students := fetcher.FetchFromGoogleSheets()
	events, projects := fetcher.FetchFromOneDriveFiles()

	bundle := bundle.BuildBundle(
		&associates,
		&assocEntries,
		&teamEntries,
		&events,
		&positions,
		&projects)

	generator.InsertGeneratedSubstitutions(&bundle, "content")

	taskStudentBundle := points.TaskStudentBundle{
		Tasks:    &tasks,
		Students: &students,
	}

	file, _ := json.Marshal(taskStudentBundle)
	_ = ioutil.WriteFile("static/points.json", file, 0644)
}
