# utmist.gitlab.io

Club website for the [University of Toronto Machine Intelligence Student Team (UTMIST)](https://utmist.gitlab.io). It's composed the following parts.

- [Google Sheets](https://developers.google.com/sheets) serve as databases.
- The **fetcher** package pulls from the databases and populates **Associate**, **Event**, and **Project** objects.
- The **generator** package uses the objects fetched to generate **markdown** content pages.
- **Hugo** generates the static site website locally or with **GitLab Pages**.

## Connections

- The fetcher/generator and Hugo are run in [GitLab CI](https://docs.gitlab.com/ce/ci/) and fed into [GitLab Pages](https://docs.gitlab.com/ce/user/project/pages/).
- The [UTMIST Assistant (MISTA)](https://gitlab.com/utmist/mista) can trigger a job to regenerate when responding to commands in our [Discord Server](https://discord.gg/88mSPw8). If MISTA is offline, jobs must be triggered manually through the [GitLab CI/CD Pipeline Manager](https://gitlab.com/utmist/utmist.gitlab.io/pipelines).

## Prerequisites

- [Go](https://golang.org/). Remember to set the GOPATH.
- [Hugo](https://github.com/gohugoio/hugo/releases), `>= 0.61`. GitLab CI uses `0.66`.
- See [Google Sheets API for Go](https://developers.google.com/sheets/api/quickstart/go).

## Details

- `.env` contains our secrets and other variables. Refer to `.env.copy` for the required variables.
  - Credentials for the Google Sheets API.
  - Google Sheet IDs and ranges.
  - Discord server invite link.
  - List of departments to show on the webiste. This will matter if department structures change or are listed differently.
    - A year might have a different list. Use the `DEPARTMENTS_YEAR` variable to specify which year you're using.
- `main.go` is the driver package and utilizes packages (including `fetcher` and `generator`) under `./src`.
- Associate/Event/Project pages are generated in generated folders `associate`/`events`/`projects` under `./contents`.
  - These files are not meant to be committed to the repository as they should be generated with fresh content each time. Be sure these are in `.gitignore`.
- `./assets` contains some data and templates files we copy and stitch content into.
  - `./assets/config.yaml => config.yaml`
  - `./assets/team.md => content/team/list.md`
  - `./assets/events.md => content/events/list.md`
  - `./assets/utsg_buildings.txt` contains the codes and numbers for buildings on the UofT St. George campus.
- `config.yaml` configures what the website looks like.
  - It defines the Docker image we use for `Go` and `Hugo`.
  - The `generator` stitches several links into a copy of `./assets/config.yaml`.

### GitLab

- We use GitLab CI and GitLab Pages to host this website.
  - GitLab CI has its own environment registry.

### Usage

- Clone the repository under `$GOPATH/gitlab.com/utmist/` and initialize theme submodule.

  ```sh
  git clone https://gitlab.com/utmist/utmist.gitlab.io.git
  git submodule update --init --recursive
  ```

- Paste the environment variables. Refer to `.env.copy` for the required variables.
- Run the `fetcher/generator` script.

  ```sh
  go run main.go
  ```

  Or if you prefer to compile first.

  ```sh
  go build
  ./utmist.gitlab.io
  ```

- Run `hugo` in debugging mode to host the website on `localhost:1313`.

  ```sh
  hugo server -D
  ```

## Developers

- **[Rupert Wu](https://leglesslamb.gitlab.io)**
- [Salim](https://msanvarov.github.io/personal-portfolio)
