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

	events, associates, projects := fetcher.Fetch()
	generator.GeneratePages(events, associates, projects)
}
