// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package dbaccess

import (
	"bytes"
	"errors"
	"strings"

	"crypto/rand"
	"crypto/sha512"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type hashFun_t func(data []byte) []byte

var db *sql.DB
var hashFun hashFun_t

const SALT_SIZE = 32

func Open(dbfile string) error {
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT version, hashfun FROM meta")
	if err != nil {
		return err
	}
	defer rows.Close()
	var version string
	var hashfun string

	if rows.Next() {
		err = rows.Scan(&version, &hashfun)
		if err != nil {
			return err
		}
		if strings.Compare(hashfun, "sha512") == 0 {
			hashFun = sha512slice
		} else {
			return errors.New("hash function unsupported")
		}
	} else {
		return errors.New("no meta table")
	}

	_, err = db.Query("SELECT user, hashpass, salt FROM users")

	return err
}

func Close() {
	db.Close()
}

func sha512slice(data []byte) []byte {
	hasharray := sha512.Sum512(data)
	return hasharray[:]
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
	res, err := stmt.Exec(user, hashFun(append(salt, pass...)), salt)
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
	_, err = stmt.Exec(hashFun(append(salt, pass...)), salt, user)
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
	if bytes.Compare(hashFun(append(salt, pass...)), hashpass) != 0 {
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
