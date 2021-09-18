package main

import (
	"flag"
	"log"

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

	associates, assocEntries, teamEntries, positions := fetcher.FetchFromGoogleSheets()
	events, projects := fetcher.FetchFromOneDriveFiles()

	bundle := bundle.BuildBundle(
		&associates,
		&assocEntries,
		&teamEntries,
		&events,
		&positions,
		&projects)

	generator.InsertGeneratedSubstitutions(&bundle, "content")
}
