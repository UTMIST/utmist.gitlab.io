# utmist.gitlab.io

Club website for the University of Toronto Machine Intelligence Student Team (UTMIST). It's composed the following parts.

- Google Sheets serve as databases.
- The **fetcher** package pulls from the databases and populates **Associate**, **Event**, and **Project** objects.
- The **generator** package uses the objects fetched to generate **markdown** content pages.
- **Hugo** generates the static site website locally or with **GitLab Pages**.

### Connections

- The fetcher/generator and Hugo are run in GitLab's CI and fed into GitLab Pages.
- The [UTMIST Assistant (MISTA)](https://gitlab.com/utmist/mista) can trigger a job to regenerate when responding to commands in our Discord Server. If MISTA is offline, jobs must be triggered manually through the [GitLab CI/CD Pipeline Manager](https://gitlab.com/utmist/utmist.gitlab.io/pipelines).

## Prerequisites

- [Go](https://golang.org/).
- [Hugo](https://github.com/gohugoio/hugo/releases), `>= 0.61`.
- See [Google Sheets API for Go](https://developers.google.com/sheets/api/quickstart/go).

## Details

- `.gitlab-ci.yml` defines what the GitLab CI will do when running a pipeline. In particular, it lists the `scripts` the CI will run, and where to look for the static site files (currently in `./public`).
- We originally utilized `credentials.json` and `token.json` as the Google Sheets API documentation had suggested. However, locally, we now just use the `.env` (similar to `.env.copy`) provided by the team workspace.
- Associate, Event, and Project pages aren't meant to be remain in the codebase; they are to be generated and used only in local testing and by GitLab CI to push to GitLab Pages.

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
- Add the `hugo-fresh` theme.
  ```
  git submodule update --init --recursive
  ```

## Usage

- `go main.go` will generate the site content using `fetcher` and `generator`.
  - `fetcher` uses credentials from `.env` to create a `token.json` for Google Sheets API access locally. GitLab CI uses environment variables.
  - It will then generate the **markdown** pages stored in `./content`.
- `hugo server -D` will run the website on `localhost:1313`.

## Rationale

- This new website [utmist.gitlab.io](https://utmist.gitlab.io) is intended to replace [utmist.github.io](utmist.github.io).
- Instead of having `travis` rebuild the website on GH pages every 24h, we will instead move towards a **Discord/Slack bot**, allowing any member of the UTMIST Workspace to run the GitLab CI using the most recent data at will.

## Developers

- **[Rupert Wu](https://leglesslamb.gitlab.io)**
- [Salim](https://msanvarov.github.io/personal-portfolio)
