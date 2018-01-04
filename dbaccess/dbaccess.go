// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package dbaccess

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func Open() {
	var err error
	db, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatalf("failed to open database %v", err)
	}
}

func Close() {
	db.Close()
}
