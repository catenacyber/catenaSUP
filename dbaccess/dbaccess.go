// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package dbaccess

import (
	"bytes"
	"errors"

	"crypto/rand"
	"crypto/sha512"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const SALT_SIZE = 32

func Open(dbfile string) error {
	var err error
	db, err = sql.Open("sqlite3", dbfile)
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

func CheckUserPass(user string, pass string) (error, uint64) {
	stmt, err := db.Prepare("SELECT rowid, hashpass, salt FROM users WHERE user = ?")
	if err != nil {
		return err, 0
	}
	var hashpass []byte
	var salt []byte
	var id uint64
	err = stmt.QueryRow(user).Scan(&id, &hashpass, &salt)
	if err != nil {
		return err, 0
	}
	hashpass1 := sha512.Sum512(append(salt, pass...))
	if bytes.Compare(hashpass1[:], hashpass) != 0 {
		err = errors.New("password does not match")
	}
	return err, id
}

func DeleteUser(user string) error {
	stmt, err := db.Prepare("DELETE FROM users WHERE user = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user)
	return err
}
