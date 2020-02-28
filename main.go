package main

import (
	fetcher "gitlab.com/utmist/utmist.gitlab.io/src/fetcher"
	generator "gitlab.com/utmist/utmist.gitlab.io/src/generator"
)

func main() {
	events, execs, projects := fetcher.Fetch()
	generator.GeneratePages(events, execs, projects)
}
