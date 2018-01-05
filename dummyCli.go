// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client for catenaSUP

package main

import (
	"log"

	clientSUP "catena/csup/client"
)

func main() {
	clientSUP.Open()
	err, id := clientSUP.AddUser("bob", "love2alice")
	if err != nil {
		log.Fatalf("error adding user: %v", err)
	}
	log.Printf("user added with id %d", id)
	err = clientSUP.ChangePass("bob", "love2oscar")
	if err != nil {
		log.Fatalf("error changing password: %v", err)
	}

	err, id = clientSUP.AddUser("alice", "wonderland")
	if err != nil {
		log.Fatalf("error adding user: %v", err)
	}
	log.Printf("user added with id %d", id)

	err, id = clientSUP.CheckUserPass("bob", "love2alice")
	if err != nil {
		log.Printf("access denied for user password: %v", err)
	} else {
		log.Printf("access granted with old password")
	}
	err, id = clientSUP.CheckUserPass("bob", "secondtry")
	if err != nil {
		log.Printf("access denied for user password: %v", err)
	} else {
		log.Printf("access granted with wrong password")
	}
	err, id = clientSUP.CheckUserPass("bob", "love2oscar")
	if err != nil {
		log.Printf("access denied with good password: %v", err)
	} else {
		log.Printf("access granted with ok password")
	}

	err = clientSUP.DeleteUser("bob")
	if err != nil {
		log.Fatalf("error deleting user: %v", err)
	}
	err, id = clientSUP.CheckUserPass("bob", "love2oscar")
	if err != nil {
		log.Printf("access denied for user password: %v", err)
	} else {
		log.Printf("access granted with deleted user")
	}
	clientSUP.Close()
}
