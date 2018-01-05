// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package main

import (
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"catena/csup/dbaccess"
	pb "catena/csup/msg"
)

const (
	port = ":5455"
)

// serverSUP implements generated CatenaUserPass
type serverSUP struct{}

// adding a user with his password
func (s *serverSUP) AddUser(ctx context.Context, in *pb.UserPass) (*pb.Id, error) {
	err, id := dbaccess.AddUser(in.User, in.Pass)
	return &pb.Id{Id: id}, err
}

// changes the password of a user
func (s *serverSUP) ChangePass(ctx context.Context, in *pb.UserPass) (*pb.Empty, error) {
	err := dbaccess.ChangePass(in.User, in.Pass)
	return &pb.Empty{}, err
}

// Checks if a user password pair is valid
func (s *serverSUP) CheckUserPass(ctx context.Context, in *pb.UserPass) (*pb.Id, error) {
	err, id := dbaccess.CheckUserPass(in.User, in.Pass)
	return &pb.Id{Id: id}, err
}

// deletes a user
func (s *serverSUP) DeleteUser(ctx context.Context, in *pb.User) (*pb.Empty, error) {
	err := dbaccess.DeleteUser(in.User)
	return &pb.Empty{}, err
}

func main() {
	if dbaccess.Open() != nil {
		log.Fatalf("failed to open database")
	}
	defer dbaccess.Close()
	//TODO port as parameter
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//TODO use SSL
	pb.RegisterCatenaUserPassServer(s, &serverSUP{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
