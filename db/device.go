package db

type Device struct {
	Id int64
	NewDevice
}

type NewDevice struct {
	Name        string
	Description string
	Token       string
}
