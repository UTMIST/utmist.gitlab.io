package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func getSeriesString(directory, seriesName, listPref string) string {
	seriesListStr := fmt.Sprintf("### %s: ", listPref)

	dirSteps := strings.Split(directory, "/")
	folderStr := dirSteps[len(dirSteps)-1]
	currentPageStr := strings.Join(strings.Split(folderStr, "-")[1:], "-")

	files, err := ioutil.ReadDir(fmt.Sprintf("%s%s", directory, "/.."))
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		labels := strings.Split(f.Name(), "-")
		if labels[0] != seriesName {
			continue
		}
		pageLabel := strings.Join(labels[1:], "-")
		if pageLabel == currentPageStr {
			seriesListStr = fmt.Sprintf("%s**%s** | ",
				seriesListStr,
				currentPageStr)
		} else {
			seriesListStr = fmt.Sprintf("%s[%s](../%s) | ",
				seriesListStr,
				pageLabel,
				strings.ToLower(f.Name()))
		}

	}
	return seriesListStr[:len(seriesListStr)-3]
}
