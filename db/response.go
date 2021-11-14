package db

type Response struct {
	Id int64
	NewResponse
}

type NewResponse struct {
	TaskId int64
	State  int16
}
