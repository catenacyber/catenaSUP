// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client for catenaSUP

package client

import (
	"errors"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "catena/csup/msg"
)

const (
	address = "localhost:5455"
)

var conn *grpc.ClientConn
var rpcConn pb.CatenaUserPassClient

func Open() {
	//TODO package avec init, clean, et wrapper adduser et cie
	var err error
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	rpcConn = pb.NewCatenaUserPassClient(conn)
}

func Close() {
	conn.Close()
}

func AddUser(user string, pass string) error {
	r, err := rpcConn.AddUser(context.Background(), &pb.UserPass{User: user, Pass: pass})
	if err != nil {
		log.Printf("could not add user: %v", err)
		return err
	}
	if r.Result != pb.Status_SUCCESS {
		log.Printf("Failed adding user")
		return errors.New("Server failed adding user")
	}
	return nil
}
