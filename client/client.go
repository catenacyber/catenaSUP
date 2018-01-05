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

var conn *grpc.ClientConn
var rpcConn pb.CatenaUserPassClient

func Open(address string, tlsOptions []grpc.DialOption) {
	var err error
	conn, err = grpc.Dial(address, tlsOptions...)
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

func CheckUserPass(user string, pass string) (error, uint64) {
	id, err := rpcConn.CheckUserPass(context.Background(), &pb.UserPass{User: user, Pass: pass})
	if err != nil {
		return err, 0
	}
	return nil, id.Id
}

func DeleteUser(user string) error {
	_, err := rpcConn.DeleteUser(context.Background(), &pb.User{User: user})
	return err
}
