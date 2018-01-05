// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client for catenaSUP

package client

import (
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

func AddUser(user string, pass string) (error, uint64) {
	id, err := rpcConn.AddUser(context.Background(), &pb.UserPass{User: user, Pass: pass})
	if err != nil {
		log.Printf("could not add user: %v", err)
		return err, 0
	}
	return nil, id.Id
}

func ChangePass(user string, pass string) error {
	_, err := rpcConn.ChangePass(context.Background(), &pb.UserPass{User: user, Pass: pass})
	return err
}
