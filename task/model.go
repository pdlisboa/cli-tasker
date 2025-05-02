package task

type Task struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
