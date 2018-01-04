// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "catena/csup/msg"
)

const (
	port = ":5455"
)

// serverSUP implements generated CatenaUserPass
type serverSUP struct{}

// adding a user with his password
func (s *serverSUP) AddUser(ctx context.Context, in *pb.UserPass) (*pb.Status, error) {
	//TODO
	return &pb.Status{Result: *pb.Status_SUCCESS}, nil
}

// changes the password of a user
func (s *serverSUP) ChangePass(ctx context.Context, in *pb.UserPass) (*pb.Status, error) {
	//TODO
	return &pb.Status{Result: *pb.Status_SUCCESS}, nil
}

// Checks if a user password pair is valid
func (s *serverSUP) CheckUserPass(ctx context.Context, in *pb.UserPass) (*pb.Status, error) {
	//TODO
	return &pb.Status{Result: *pb.Status_SUCCESS}, nil
}

// deletes a user
func (s *serverSUP) DeleteUser(ctx context.Context, in *pb.User) (*pb.Status, error) {
	//TODO
	return &pb.Status{Result: *pb.Status_SUCCESS}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCatenaUserPassServer(s, &serverSUP{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
