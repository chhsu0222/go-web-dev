package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB is the connection object.
var DB *sql.DB

func init() {
	// Initialize connection string.
	connStr := "postgres://chhsu:example@localhost/localTest?sslmode=disable"

	// Initialize connection object.
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Println("You connected to your database.")

	// Create table.
	createTableStr := `CREATE TABLE IF NOT EXISTS books (
		isbn    char(14)     PRIMARY KEY NOT NULL,
		title   varchar(255) NOT NULL,
		author  varchar(255) NOT NULL,
		price   decimal(5,2) NOT NULL
	  );`

	_, err = DB.Exec(createTableStr)
	if err != nil {
		log.Panic(err)
	}
}
