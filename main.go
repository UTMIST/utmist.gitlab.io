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

	events, associates, positions, projects, pastProjects := fetcher.Fetch()
	generator.GeneratePages(&events, &associates, &positions, &projects, &pastProjects)
	generator.GenerateConfig(&events, &projects)
}
