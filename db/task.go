package db

type Task struct {
	Id int64
	NewTask
}

type NewTask struct {
	Title  string
	Device int64
	State  int16
}
