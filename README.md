# utmist.gitlab.io

Club website for the [University of Toronto Machine Intelligence Student Team (UTMIST)](https://utmist.gitlab.io).

[![pipeline status](https://gitlab.com/utmist/utmist.gitlab.io/badges/master/pipeline.svg)](https://gitlab.com/utmist/utmist.gitlab.io/-/commits/master)

## Overview

- [Google Sheets](https://developers.google.com/sheets) and [Microsoft OneDrive](https://onedrive.live.com/) serve as databases.
- The `fetcher` package pulls from [Google Sheets](https://developers.google.com/sheets) and populates **Associate/Department** objects.
- The `generator` package uses the objects fetched to generate **markdown** content pages.
- `onedeath downloads a folder of custom pages (written manually by club associates) from [Microsoft OneDrive](https://onedrive.live.com/).
- **Hugo** generates the static site website locally or with **GitLab Pages**.

### Connections

- The `fetcher`/`generator`, `onedeath` and Hugo are run in [GitLab CI](https://docs.gitlab.com/ce/ci/) and fed into [GitLab Pages](https://docs.gitlab.com/ce/user/project/pages/).
- The [UTMIST Runner (MISTR)](https://gitlab.com/utmist/mistr) can trigger a job to regenerate when responding to commands in our [Discord Server](https://discord.gg/88mSPw8). If MISTR is offline, jobs must be triggered manually through the [GitLab CI/CD Pipeline Manager](https://gitlab.com/utmist/utmist.gitlab.io/pipelines).

### Prerequisites

- [Go](https://golang.org/). Put this project in `$GOPATH/utmist/`.
- [Hugo](https://github.com/gohugoio/hugo/releases), `>= 0.61`.
- [Lua](https://www.lua.org/).
- [wget](https://www.gnu.org/software/wget/).

### Dependencies

- [godotenv](https://pkg.go.dev/github.com/joho/godotenv)
- [Google Sheets API for Go](https://pkg.go.dev/google.golang.org/api)

## Details

Full details can be found on [our Wiki](https://gitlab.com/utmist/utmist.gitlab.io/-/wikis).

- [Content Management](https://gitlab.com/utmist/utmist.gitlab.io/-/wikis/Exec-Team-Guide-To-Content-Management)
- Development (Coming Soon)

### GitLab

- We use GitLab CI and GitLab Pages to host this website.
  - GitLab CI has its own environment registry.

### Usage

- Clone the repository under `$GOPATH/gitlab.com/utmist/`.

  ```sh
  cd $GOPATH/gitlab.com/utmist/
  git clone https://gitlab.com/utmist/utmist.gitlab.io.git
  cd utmist.gitlab.io
  ```

- Initialize the Hugo ReFresh and OneDeath submodules.

  ```sh
  make dep
  ```

- Paste the environment variables into `.env`. Refer to `.env.copy` for the required variables.
- Run the `fetcher/generator` script.

  ```sh
  make full
  ```

- Run `hugo` in future mode (include pages dated in the future) to host the website on `localhost:1313`.

  ```sh
  hugo server -F
  ```

## Development

- [**GitLab**](https://gitlab.com/utmist/utmist.gitlab.io)
  - Working issues and Merge Requests (MRs) are reviewed.
  - Bug reports and feature requests are preferred.
- [**GitHub (Mirror)**](https://github.com/utmist/utmist.gitlab.io)
  - Bug reports and feature requests are accepted.
- This project is maintained by the [Infrastructure Department at UTMIST](https://utmist.gitlab.io/team/infrastructure).
- If you’re a member of UTMIST and would like to contribute or learn development through this project, you can join our [Discord Server](https://discord.gg/88mSPw8) and let us know in #infrastructure.

## Acknowledgements

- [Salim Anvarov](https://msanvarov.github.io/personal-portfolio) for advising on Docker and Go Modules.
- [Lingkai (Rain) Shen](https://www.linkedin.com/in/lingkai-shen/) for building [utmist.github.io](https://github.com/utmist/utmist.github.io) and advising on Google Sheets.
