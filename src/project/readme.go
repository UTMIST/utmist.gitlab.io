package project

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// DownloadReadMe downloads the file from given URL
func DownloadReadMe(url string) []string {

	resp, err := http.Get(url)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to retrieve README file from %s: ",
			url))
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(body), "\n")

	for i := len(lines) - 1; i >= 0; i-- {
		if len(lines[i]) > 2 &&
			lines[i][0] == '#' &&
			lines[i][1] != '#' {
			lines[i] = ""
		}
	}

	return lines
}
