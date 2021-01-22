package fetcher

import (
	"fmt"
	"os"
	"strconv"

	"gitlab.com/utmist/utmist.gitlab.io/src/points"
	"google.golang.org/api/sheets/v4"
)

func fetchTasksAndStudents(
	srv *sheets.Service,
	firstYear,
	lastYear int) (map[string][]string, map[string]map[string]points.Student) {

	students := map[string]map[string]points.Student{}
	tasks := map[string][]string{}
	sheetID := os.Getenv("STUDENTS_SHEET_ID")
	for y := lastYear; y >= firstYear; y-- {
		yearStudents := map[string]points.Student{}

		indexRange := fmt.Sprintf("%d!%s", y, os.Getenv("STUDENTS_SHEET_INDEX"))
		resp, err := fetchValues(
			srv,
			fmt.Sprintf("Students (%d)", y),
			sheetID,
			indexRange)

		if err != nil {
			continue
		}

		index := resp.Values

		sheetRange := fmt.Sprintf("%d!%s", y, os.Getenv("STUDENTS_SHEET_RANGE"))
		resp, _ = fetchValues(
			srv,
			fmt.Sprintf("Students (%d)", y),
			sheetID,
			sheetRange)

		// find number of tasks from first row
		numTasks, yearTasks := points.LoadTasks(resp.Values[0])
		for i, row := range resp.Values[1:] {
			curStudent := points.LoadStudent(index[i][0].(string), row, numTasks)
			yearStudents[curStudent.Username] = curStudent
		}
		yearStr := fmt.Sprintf("%s-%s", strconv.Itoa(y-1), strconv.Itoa(y))
		students[yearStr] = yearStudents
		tasks[yearStr] = yearTasks
	}

	return tasks, students
}
