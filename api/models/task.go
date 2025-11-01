package models

type ScheduleTaskRequest struct {
	TaskId string   `json:"taskId"`
	Dates  []string `json:"dates"`
}

type TaskScheduleItem struct {
	TaskId string `json:"taskId"`
	Date   string `json:"date"`
}
