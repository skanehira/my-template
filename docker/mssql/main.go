package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

var server = "localhost"
var port = 1433
var user = "sa"
var password = "mssqlPassword!"
var database = "test"

func main() {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from users where id = @id")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(sql.Named("id", 1))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id       int
			name     string
			birthday time.Time
		)

		if err := rows.Scan(&id, &name, &birthday); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("id :%d, name: %s, bithday: %v", id, name, birthday)
	}

}
