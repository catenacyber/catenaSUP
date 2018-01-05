// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client for catenaSUP

package main

import (
	"log"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatalf("failed to open database %v", err)
	}
	stmt, err := db.Prepare("CREATE TABLE users (user TEXT UNIQUE, hashpass BLOB, salt BLOB)")
	if err != nil {
		log.Fatalf("failed to prepare create request %v", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("failed to execute create request %v", err)
	}
	//TODO metadata table
}
