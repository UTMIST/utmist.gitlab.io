# utmist.gitlab.io

Club website for the University of Toronto Machine Intelligence Student Team (UTMIST)

## Prerequisites

- [Go](https://golang.org/).
- [Hugo](https://github.com/gohugoio/hugo/releases), `>= 0.61`.

## Setup/Housekeeping

- `go-get.sh` downloads all the required Go packages.
- `update-fresh.sh` will refresh the `hugo-fresh` theme.
- `.gitlab-ci.yml` defines what the GitLab CI will do when running a pipeline.

## Usage

- `go main.go` will (in the near future) generate the site content using `generator`.
- `hugo server -D` will run the website on `localhost:1313`.
