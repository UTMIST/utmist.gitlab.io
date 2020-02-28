package main

import (
	fetcher "gitlab.com/utmist/utmist.gitlab.io/src/fetcher"
	generator "gitlab.com/utmist/utmist.gitlab.io/src/generator"
)

func main() {
	events, associates, projects := fetcher.Fetch()
	generator.GeneratePages(events, associates, projects)
}
