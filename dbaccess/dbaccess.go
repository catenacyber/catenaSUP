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
	"github.com/lhecker/argon2"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

//type definition of the hashing function
type hashFun_t func(data []byte, salt []byte) (error, []byte)

//global variable : database object connection
var db *sql.DB

//global varibale : current hashing function
var hashFun hashFun_t

//contant parameter : size of salt in bytes
const SALT_SIZE = 32

func Open(dbfile string) error {
	var err error
	db, err = sql.Open("sqlite3", dbfile)
	if err != nil {
		return err
	}

	//consistency checks of database
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
		if strings.Compare(hashfun, "argon2") == 0 {
			hashFun = argon2slice
		} else if strings.Compare(hashfun, "sha512") == 0 {
			hashFun = sha512slice
		} else {
			return errors.New("hash function unsupported")
		}
	} else {
		return errors.New("no meta table")
	}

	_, err = db.Query("SELECT user, hashpass, salt FROM users")
	//could further check types of columns

	return err
}

func Close() {
	db.Close()
}

func sha512slice(data []byte, salt []byte) (error, []byte) {
	//sha512 function with slices and not array as output
	hasharray := sha512.Sum512(append(salt, data...))
	return nil, hasharray[:]
}

func argon2slice(data []byte, salt []byte) (error, []byte) {
	//sha512 function with slices and not array as output
	//TODO
	cfg := argon2.DefaultConfig()
	cfg.SaltLength = uint32(len(salt))
	raw, err := cfg.Hash(data, salt)
	return err, raw.Hash
}

//top level db access functions
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
	err, hashed := hashFun([]byte(pass), salt)
	if err != nil {
		return err, 0
	}
memzerostr(&pass)
	res, err := stmt.Exec(user, hashed, salt)
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
	err, hashed := hashFun([]byte(pass), salt)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(hashed, salt, user)
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
	//could check hashpass length for upgrading hash function
	err, hashed := hashFun([]byte(pass), salt)
	if err != nil {
		return err, 0
	}
	if bytes.Compare(hashed, hashpass) != 0 {
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
