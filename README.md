# utmist.gitlab.io

Club website for the University of Toronto Machine Intelligence Student Team (UTMIST). It's built in **Go**, pulls our "database" sheets from the Google Sheets API, and generates **markdown** content pages. We then use **Hugo** to generate the static site with these pages on the **GitLab CI**.

## Prerequisites

- [Go](https://golang.org/).
- [Hugo](https://github.com/gohugoio/hugo/releases), `>= 0.61`.
- See [Google Sheets API for Go](https://developers.google.com/sheets/api/quickstart/go).

## Setup/Housekeeping

- Clone into the `GOPATH` using `SSH` or `HTTPS`.
  ```
  cd $GOPATH/src/gitlab.com/utmist
  ```
  `SSH`
  ```
  git clone git@gitlab.com:utmist/utmist.gitlab.io.git
  ```
  `HTTPS`
  ```
  git clone https://gitlab.com/utmist/utmist.gitlab.io.git
  ```

* Add the `hugo-fresh` theme.
  ```
  git submodule update --init --recursive
  ```
* `.gitlab-ci.yml` defines what the GitLab CI will do when running a pipeline. In particular, it lists the `scripts` the CI will run, and where to look for the static site files (currently in `./public`).
* We originally utilized `credentials.json` and `token.json` as the Google Sheets API documentation had suggested. However, locally, we now just use the `.env` (similar to `.env.copy`) provided by the team workspace.

## Usage

- `go main.go` will generate the site content using `generator`.
  - This consists a `fetch` script to pull data from the Google Sheets, and will use `credentials.json` to generate a `token.json` (looking similar to `token.copy.json`) if this file doesn't exist.
  - It will then generate the **markdown** pages stored in `./content`.
- `hugo server -D` will run the website on `localhost:1313`.

## Development & Planning

- This new website [utmist.gitlab.io](https://utmist.gitlab.io) is intended to replace [utmist.github.io](utmist.github.io).
- Instead of having `travis` rebuild the website on GH pages every 24h, we will instead move towards a **Discord/Slack bot**, allowing any member of the UTMIST Workspace to run the GitLab CI using the most recent data at will.

## Developers

- [Rupert Wu](https://leglesslamb.gitlab.io)
- [Salim](https://msanvarov.github.io/personal-portfolio)
