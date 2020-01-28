# utmist.gitlab.io

Club website for the University of Toronto Machine Intelligence Student Team (UTMIST). It's built in **Go**, pulls our "database" sheets from the Google Sheets API, and generates **markdown** content pages. We then use **Hugo** to generate the static site with these pages on the **GitLab CI**.

## Prerequisites

- [Go](https://golang.org/).
- [Hugo](https://github.com/gohugoio/hugo/releases), `>= 0.61`.
- See [Google Sheets API for Go](https://developers.google.com/sheets/api/quickstart/go).

## Setup/Housekeeping

- `go-get.sh` downloads all the required Go packages.
- `update-fresh.sh` will refresh the `hugo-fresh` theme.
- `.gitlab-ci.yml` defines what the GitLab CI will do when running a pipeline. In particular, it lists the `scripts` the CI will run, and where to look for the static site files (currently in `./public`).
- Get `credentials.json` from signing up on the Google Could Platform. It should have a similar form to `credentials.copy.json`.
- You'll need access to the UTMIST drive folders for your `credentials.json` to work. This is usually done by logging into the **VP Communications** account; you might also be able to share the UTMIST GDrive folders with your personal account and generate `credentials.json` yourself.
- Get the sheet IDs and ranges and put them in `.env` (similar to `.env.copy`).

## Usage

- `go main.go` will generate the site content using `generator`.
  - This consists a `fetch` script to pull data from the Google Sheets, and will use `credentials.json` to generate a `token.json` (looking similar to `token.copy.json`) if this file doesn't exist.
  - It will then generate the **markdown** pages stored in `./content`.
- `hugo server -D` will run the website on `localhost:1313`.

## Development & Planning

- This new website [utmist.gitlab.io](https://utmist.gitlab.io) is intended to replace [utmist.github.io](utmist.github.io).
- Instead of having `travis` rebuild the website on GH pages every 24h, we will instead move towards a **Slack bot**, allowing any member of the UTMIST Workspace to run the GitLab CI using the most recent data at will.
