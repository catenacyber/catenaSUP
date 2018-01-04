// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package dbaccess

import (
	"crypto/rand"
	"crypto/sha512"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const SALT_SIZE = 32

func Open() error {
	var err error
	//TODO name as parameter
	db, err = sql.Open("sqlite3", "./foo.db")
	//TODO further checks if db is ok
	return err
}

func Close() {
	db.Close()
}

func AddUser(user string, pass string) error {
	stmt, err := db.Prepare("INSERT INTO users VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	salt := make([]byte, SALT_SIZE)
	_, err = rand.Read(salt)
	if err != nil {
		return err
	}
	hashpass := sha512.Sum512(append(salt, pass...))
	_, err = stmt.Exec(user, hashpass[:], salt)
	if err != nil {
		return err
	}
	return nil
}
