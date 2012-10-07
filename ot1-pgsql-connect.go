// Download the pgsql library from github
//	 go help get
//	 go get github.com/lxn/go-pgsql
package main

import (
	"database/sql"
	"fmt"
	"log"
)

import (
	_ "github.com/lxn/go-pgsql"
)

func main() {

	// connect to database
	db, err := sql.Open("postgres", "user=chad password= host=localhost port=5432 dbname=temperature sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var msg string

	err = db.QueryRow("SELECT $1 || ' ' || $2;", "Hello", "SQL").Scan(&msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
