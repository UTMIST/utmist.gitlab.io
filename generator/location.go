package generator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const utsgBuildingsFile = "assets/utsg_buildings.txt"
const utsgMapLinkBase = "http://map.utoronto.ca/utsg/building"

// Building represents some building at UofT.
type Building struct {
	code   string
	name   string
	number string
}

// Get list of UofT building structs.
func getUofTBuildingsList() (map[string]Building, error) {

	// Load the UTSG building file.
	utsgBuildingsFile, err := os.Open(utsgBuildingsFile)
	if err != nil {
		return nil, err
	}

	buildings := map[string]Building{}

	// Read a building per line, adding the building to the list.
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

// Return building link from map.utoronto.ca
func (b *Building) getUofTMapsLink(room string) string {
	return fmt.Sprintf("[%s](%s/%s) %s",
		b.name,
		utsgMapLinkBase,
		b.number,
		room)
}
