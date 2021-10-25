package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Options struct {
	Driver string
	User   string
	Pass   string
	Addr   string
	DBName string
}

func Init(opts Options) {
	switch opts.Driver {
	case "mysql":
		cfg := mysql.Config{
			User:   opts.User,
			Passwd: opts.Pass,
			Net:    "tcp",
			Addr:   opts.Addr,
			DBName: opts.DBName,
		}
		// Get a database handle.
		var err error
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatal(err)
		}
	default:
		log.Fatalf("Don't know how to init a DB of type %s", opts.Driver)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// Inserts a user, returning the assigned ID
func InsertUser(u User) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, email, admin, token) VALUES (?, ?, ?, ?)", u.Name, u.Email, u.Admin, u.Token)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Deletes a user, returning an error if it fails
func DeleteUser(id int64, mustExist bool) error {
	result, err := db.Exec("DELETE FROM users U WHERE U.id = ?", id)
	if mustExist {
		if affected, _ := result.RowsAffected(); affected == 0 {
			return errors.New("Id not found")
		}
	}
	return err
}

// Inserts a device, returning the assigned ID
func InsertDevice(d Device) (int64, error) {
	result, err := db.Exec("INSERT INTO devices (name, description, token) VALUES (?, ?, ?)", d.Name, d.Description, d.Token)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Deletes a device, returning an error if it fails
func DeleteDevice(id int64, mustExist bool) error {
	result, err := db.Exec("DELETE FROM devices D WHERE D.id = ?", id)
	if mustExist {
		if affected, _ := result.RowsAffected(); affected == 0 {
			return errors.New("Id not found")
		}
	}
	return err
}

// Inserts a task, returning the assigned ID
func InsertTask(t Task) (int64, error) {
	result, err := db.Exec("INSERT INTO tasks (title, device, state) VALUES (?, ?, ?)", t.Title, t.Device, t.State)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Deletes a task, returning an error if it fails
func DeleteTask(id int64, mustExist bool) error {
	result, err := db.Exec("DELETE FROM tasks T WHERE T.id = ?", id)
	if mustExist {
		if affected, _ := result.RowsAffected(); affected == 0 {
			return errors.New("Id not found")
		}
	}
	return err
}

// Inserts a response for a task, returning the assigned ID
func InsertResponse(r Response) (int64, error) {
	result, err := db.Exec("INSERT INTO responses (taskid, state) VALUES (?, ?)", r.TaskId, r.State)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

// Deletes a task, returning an error if it fails
func DeleteResponse(id int64, mustExist bool) error {
	result, err := db.Exec("DELETE FROM tasks T WHERE T.id = ?", id)
	if mustExist {
		if affected, _ := result.RowsAffected(); affected == 0 {
			return errors.New("Id not found")
		}
	}
	return err
}

func GrantAccess(u User, d Device) {
	db.Exec("INSERT INTO controls (userid, deviceid) VALUES (?, ?)", u.Id, d.Id)
}

func RevokeAccess(u User, d Device) {
	db.Exec("DELETE FROM controls C WHERE C.userid = ? AND C.deviceid = ?", u.Id, d.Id)
}
