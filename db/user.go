package db

type User struct {
	Id     int64
	Name   string
	Email  string
	Admin  bool
	Token  string
	Active bool
}
