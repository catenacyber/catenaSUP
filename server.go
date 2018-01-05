// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Server for catenaSUP

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"catena/csup/dbaccess"
	pb "catena/csup/msg"
)

var (
	tlsOn    = flag.Bool("tls", true, "Connection uses TLS if true")
	cliCert  = flag.String("cli_cert", "./client.crt", "The TLS certificate file for client")
	servCert = flag.String("serv_cert", "./serv.crt", "The TLS certificate file for server")
	servKey  = flag.String("serv_key", "./serv.key", "The TLS key file of server")
	port     = flag.Int("port", 1541, "The server port")
	dbfile   = flag.String("dbfile", "./csup.db", "The database file")
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
	flag.Parse()
	err := dbaccess.Open(*dbfile)
	if err != nil {
		log.Fatalf("failed to open database : %v", err)
	}
	defer dbaccess.Close()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var tlsOptions []grpc.ServerOption
	if *tlsOn {
		cert, err := tls.LoadX509KeyPair(*servCert, *servKey)
		if err != nil {
			log.Fatalf("failed to certificate and key: %v", err)
		}
		clientCert, err := ioutil.ReadFile(*cliCert)
		if err != nil {
			log.Fatalf("Unable to open client cert", err)
		}
		clientCertPool1 := x509.NewCertPool()
		clientCertPool1.AppendCertsFromPEM(clientCert)
		tlsConf := &tls.Config{Certificates: []tls.Certificate{cert}, ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: clientCertPool1}
		tlsOptions = []grpc.ServerOption{grpc.Creds(credentials.NewTLS(tlsConf))}
	}
	s := grpc.NewServer(tlsOptions...)
	pb.RegisterCatenaUserPassServer(s, &serverSUP{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
