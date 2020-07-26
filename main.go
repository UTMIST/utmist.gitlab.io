package main

import (
	"flag"
	"log"

	"github.com/joho/godotenv"
	"gitlab.com/utmist/utmist.gitlab.io/src/fetcher"
	"gitlab.com/utmist/utmist.gitlab.io/src/generator"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	flag.Parse()

	associates, assocEntries, events := fetcher.Fetch()

	generator.GenerateDepartmentAssociateLists(&associates, &assocEntries)
	generator.GenerateTeamDepartmentList(&associates, &assocEntries)
	generator.GenerateTeamExecutiveList(&associates, &assocEntries)
	generator.GenerateEventList(&events)
}
