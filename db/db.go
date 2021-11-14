package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

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
	log.Printf("Initiating DB connection with the following options: %v", opts)
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
		db, err = openConnection("mysql", cfg.FormatDSN())
		if err != nil {
			log.Fatalf("Couldn't open connection to database. %v\n", err)
		}
	default:
		log.Fatalf("Don't know how to init a DB of type %s", opts.Driver)
	}

	ensureConnectionOpen(db)
	fmt.Println("Connected!")
}

// Inserts a user, returning the assigned ID
func InsertUser(u CreateUserDTO) (int64, error) {
	result, err := db.Exec("INSERT INTO users (name, email, admin, token, active) VALUES (?, ?, ?, ?, ?)", u.Name, u.Email, u.Admin, u.Token, u.Active)
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

func GetUser(id int64) (User, error) {
	var user User
	if err := db.QueryRow("SELECT id, name, email, admin, token, active FROM users U WHERE U.id = ?", id).Scan(&user.Id, &user.Name, &user.Email, &user.Admin, &user.Token, &user.Active); err != nil {
		if err == sql.ErrNoRows {
			return User{}, fmt.Errorf("User with id %d not found", id)
		}
		return User{}, fmt.Errorf("GetUser failed for id %d: %v", id, err)
	}
	return user, nil
}

// Inserts a device, returning the assigned ID
func InsertDevice(d NewDevice) (int64, error) {
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
func InsertTask(t NewTask) (int64, error) {
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
func InsertResponse(r NewResponse) (int64, error) {
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

func openConnection(driver string, dsn string) (*sql.DB, error) {
	var err error
	var db *sql.DB
	log.Println("Entered openConnection")
	db, err = sql.Open(driver, dsn)
	log.Println(err)
	for i := 0; i < 6 && err != nil; i++ {
		log.Println("Connection to DB failed. Retrying in 5 seconds.")
		time.Sleep(5 * time.Second)
		db, err = sql.Open(driver, dsn)
	}
	return db, err
}

func ensureConnectionOpen(db *sql.DB) {
	var err error
	err = db.Ping()
	for i := 0; i < 12 && err != nil; i++ {
		fmt.Println("Ping failed, retrying in 5 seconds")
		time.Sleep(5 * time.Second)
		err = db.Ping()
	}
	if err != nil {
		log.Fatal(err)
	}
}
