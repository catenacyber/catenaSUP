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
		log.Printf("error adding user: %v", err)
	}
	log.Printf("user added with id %d", id)
	clientSUP.Close()
}
