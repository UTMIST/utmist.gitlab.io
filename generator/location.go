package generator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const utsgBuildingsFile = "assets/utsg_buildings.txt"

// Building represents some building at UofT.
type Building struct {
	code   string
	name   string
	number string
}

func getUofTBuildingsList() (map[string]Building, error) {
	buildings := map[string]Building{}

	utsgBuildingsFile, err := os.Open(utsgBuildingsFile)
	if err != nil {
		return buildings, err
	}

	utsgReader := bufio.NewScanner(utsgBuildingsFile)
	for utsgReader.Scan() {
		buildingParts := strings.SplitN(utsgReader.Text(), " ", 3)
		if len(buildingParts) != 3 {
			continue
		}
		bldg := Building{
			code:   buildingParts[0],
			number: buildingParts[1],
			name:   buildingParts[2],
		}
		buildings[bldg.code] = bldg
	}

	return buildings, nil
}

func (b *Building) getUofTMapsLink(room string) string {
	return fmt.Sprintf("[%s](http://map.utoronto.ca/utsg/building/%s) %s",
		b.name, b.number, room)
}
