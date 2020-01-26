package main

import (
	gen "utmist.gitlab.io/generator"
)

func main() {
	events, execs, projects := gen.Fetch()
	gen.GeneratePages(events, execs, projects)
}
