package fetcher

import (
	"os"
	"time"

	"gitlab.com/utmist/utmist.gitlab.io/src/position"
	"google.golang.org/api/sheets/v4"
)

func fetchPositions(srv *sheets.Service) []position.Position {
	positions := []position.Position{}

	sheetID := os.Getenv("POSITIONS_SHEET_ID")
	sheetRange := os.Getenv("POSITIONS_SHEET_RANGE")
	resp := fetchValues(srv, "Positions", sheetID, sheetRange)
	for _, row := range resp.Values {
		pos := position.Load(row)

		// Add the position if there's no valid deadline that passed.
		toronto, err := time.LoadLocation("America/New_York")
		if err != nil ||
			!pos.Deadline.Before(time.Now().In(toronto)) {
			positions = append(positions, pos)
		}
	}

	return positions
}
