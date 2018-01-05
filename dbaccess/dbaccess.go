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
	//TODO further checks if db is ok (has users table)
	return err
}

func Close() {
	db.Close()
}

func AddUser(user string, pass string) (error, uint64) {
	salt := make([]byte, SALT_SIZE)
	_, err := rand.Read(salt)
	if err != nil {
		return err, 0
	}
	stmt, err := db.Prepare("INSERT INTO users VALUES (?, ?, ?)")
	if err != nil {
		return err, 0
	}
	//TODO sha512 as db parameter
	hashpass := sha512.Sum512(append(salt, pass...))
	res, err := stmt.Exec(user, hashpass[:], salt)
	if err != nil {
		return err, 0
	}
	id, err := res.LastInsertId()
	return err, uint64(id)
}

func ChangePass(user string, pass string) error {
	salt := make([]byte, SALT_SIZE)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare("UPDATE users SET hashpass = ?, salt = ? WHERE user = ?")
	if err != nil {
		return err
	}
	hashpass := sha512.Sum512(append(salt, pass...))
	_, err = stmt.Exec(hashpass[:], salt, user)
	return err
}
