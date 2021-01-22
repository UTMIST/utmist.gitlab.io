package points

import (
	"crypto/sha256"
	"fmt"
)

const (
	pointsPlaceholder = "N/A"
)

func LoadTasks(data []interface{}) (int, []string) {
	tasks := []string{}
	numTasks := len(data)
	for _, task := range data {
		tasks = append(tasks, task.(string))
	}
	return numTasks, tasks
}

func LoadStudent(username string, data []interface{}, numTasks int) Student {
	pointsByTask := []string{}
	for _, taskPoints := range data {
		if taskPoints.(string) == "" {
			pointsByTask = append(pointsByTask, pointsPlaceholder)
		} else {
			pointsByTask = append(pointsByTask, taskPoints.(string))
		}
	}
	return Student{
		fmt.Sprintf("%x", sha256.Sum256([]byte(username))),
		pointsByTask,
	}
}
