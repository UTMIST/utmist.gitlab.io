package main

import (
	gen "gitlab.com/utmist/utmist.gitlab.io/generator"
)

func main() {
	events, execs, projects := gen.Fetch()
	gen.GeneratePages(events, execs, projects)
	gen.GenerateNavbarEventLinks(events)
}
