package main

import (
	"log"

	"github.com/joho/godotenv"
	fetcher "gitlab.com/utmist/utmist.gitlab.io/src/fetcher"
	generator "gitlab.com/utmist/utmist.gitlab.io/src/generator"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	assocs, descs, events, positions, pastProjs, projs := fetcher.Fetch()
	generator.GeneratePages(
		&assocs,
		&descs,
		&events,
		&positions,
		&pastProjs,
		&projs)
	generator.GenerateConfig(&events, &projs)
}
