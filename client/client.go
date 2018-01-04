// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client for catenaSUP

package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "catena/csup/msg"
)

const (
	address = "localhost:5455"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCatenaUserPassClient(conn)

	r, err := c.AddUser(context.Background(), &pb.UserPass{User: "bob", Pass: "lovealice"})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("Result: %d", r.Result)
}
