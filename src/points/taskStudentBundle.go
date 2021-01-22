package points

type TaskStudentBundle struct {
	Tasks    *map[string][]string
	Students *map[string]map[string]Student
}
