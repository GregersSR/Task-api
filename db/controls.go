package db

// Represents the "controls" relationship between users and devices
type Controls struct {
	UserId   int64
	DeviceId int64
}
