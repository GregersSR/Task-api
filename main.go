package main

import (
	"flag"
	"log"
	"os"

	"github.com/GregersSR/taskinator/db"
	"github.com/GregersSR/taskinator/restapi"
)

func main() {
	dbms := flag.String("db-driver", "mysql", "Override the DB driver to use")
	dbaddr := flag.String("db", "db", "Override the location of the DBMS to use for persistent storage")
	dbname := flag.String("dbname", "taskinator", "Override the name of the database to use for persistent storage")
	flag.Parse()
	dbuser := os.Getenv("DBUSER")
	dbpass := os.Getenv("DBPASS")
	log.Println("Initiating DB connection")
	db.Init(db.Options{
		Driver: *dbms,
		User:   dbuser,
		Pass:   dbpass,
		Addr:   *dbaddr,
		DBName: *dbname,
	})
	log.Println("Initiated DB connection")
	restapi.Serve()
}
