// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client library for catenaSUP

package client

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/catena/CatenaSUP/msg"
)

//global variable : connection objects
var conn *grpc.ClientConn
var rpcConn pb.CatenaUserPassClient

func Open(address string, tlsOptions []grpc.DialOption) error {
	var err error
	conn, err = grpc.Dial(address, tlsOptions...)
	if err != nil {
		return err
	}
	rpcConn = pb.NewCatenaUserPassClient(conn)
	return nil
}

func Close() {
	conn.Close()
}

//RPC stubs
func AddUser(user string, pass string) (error, uint64) {
	id, err := rpcConn.AddUser(context.Background(), &pb.UserPass{User: user, Password: pass})
	if err != nil {
		return err, 0
	}
	return nil, id.Id
}

func ChangePass(user string, pass string) error {
	_, err := rpcConn.ChangePass(context.Background(), &pb.UserPass{User: user, Password: pass})
	return err
}

func CheckUserPass(user string, pass string) (error, uint64) {
	id, err := rpcConn.CheckUserPass(context.Background(), &pb.UserPass{User: user, Password: pass})
	if err != nil {
		return err, 0
	}
	return nil, id.Id
}

func DeleteUser(user string) error {
	_, err := rpcConn.DeleteUser(context.Background(), &pb.User{User: user})
	return err
}
