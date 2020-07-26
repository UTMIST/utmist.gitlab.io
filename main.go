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

	associates, entries := fetcher.Fetch()

	generator.GenerateDepartmentAssociateLists(&associates, &entries)
	generator.GenerateTeamExecutiveList(&associates, &entries)

}
