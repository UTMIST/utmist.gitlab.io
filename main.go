package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	fetcher "gitlab.com/utmist/utmist.gitlab.io/src/fetcher"
	generator "gitlab.com/utmist/utmist.gitlab.io/src/generator"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	deptsFlagPtr := flag.Bool("depts", false, "Whether to Generate Department Pages")
	eventsFlagPtr := flag.Bool("events", false, "Whether to Generate Event Pages")
	flag.Parse()

	assocs, descs, events, positions, pastProjs, projs := fetcher.Fetch()
	generator.GeneratePages(
		&assocs,
		&descs,
		&events,
		&positions,
		&pastProjs,
		&projs,
		*deptsFlagPtr,
		*eventsFlagPtr)
	generator.GenerateConfig(&events, &projs)
}
