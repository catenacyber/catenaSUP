// Copyright (c) 2018 Catena cyber
// Author Philippe Antoine <p.antoine@catenacyber.fr>
// Go client example for catenaSUP

package main

import (
	"flag"
	"fmt"
	"log"

	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	clientSUP "github.com/catena/CatenaSUP/client"
)

var (
	tlsOn    = flag.Bool("tls", true, "Connection uses TLS if true")
	cliKey   = flag.String("cli_key", "./client.key", "The TLS key file of client")
	cliCert  = flag.String("cli_cert", "./client.crt", "The TLS certificate file for client")
	servCert = flag.String("serv_cert", "./serv.crt", "The TLS certificate file for server")
	server   = flag.String("server", "localhost:1541", "Server address")
)

func main() {
	//parse options
	flag.Parse()

	//prepares tls connection
	var tlsOptions []grpc.DialOption
	if *tlsOn {
		cert, err := tls.LoadX509KeyPair(*cliCert, *cliKey)
		if err != nil {
			log.Fatalf("failed to certificate and key: %v", err)
		}
		serverCert, err := ioutil.ReadFile(*servCert)
		if err != nil {
			log.Fatalf("Unable to open server cert", err)
		}
		serverCertPool1 := x509.NewCertPool()
		block, _ := pem.Decode(serverCert)
		if block == nil {
			log.Fatalf("failed to parse certificate")
		}
		xCert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			log.Fatalf("failed to parse certificate: %v", err)
		}
		serverCertPool1.AddCert(xCert)
		tlsConf := &tls.Config{Certificates: []tls.Certificate{cert}, RootCAs: serverCertPool1, ServerName: xCert.Subject.CommonName}
		tlsOptions = append(tlsOptions, grpc.WithTransportCredentials(credentials.NewTLS(tlsConf)))
	} else {
		tlsOptions = append(tlsOptions, grpc.WithInsecure())
	}

	//open connection
	err := clientSUP.Open(*server, tlsOptions)
	if err != nil {
		log.Fatalf("failed to open connection: %v", err)
	}
	defer clientSUP.Close()

	//some dummy operations for testing :
	//add user, change his password, adds another user, checks user-pass and deletes user
	err, id := clientSUP.AddUser("bob", "love2alice")
	if err != nil {
		log.Fatalf("error adding user: %v", err)
	}
	fmt.Printf("user added with id %d\n", id)
	err = clientSUP.ChangePass("bob", "love2oscar")
	if err != nil {
		log.Fatalf("error changing password: %v", err)
	}

	err, id = clientSUP.AddUser("alice", "wonderland")
	if err != nil {
		log.Fatalf("error adding user: %v", err)
	}
	fmt.Printf("user added with id %d\n", id)

	err, id = clientSUP.CheckUserPass("bob", "love2alice")
	if err != nil {
		fmt.Printf("access denied for user password: %v\n", err)
	} else {
		fmt.Printf("access granted with old password\n")
	}
	err, id = clientSUP.CheckUserPass("bob", "secondtry")
	if err != nil {
		fmt.Printf("access denied for user password: %v\n", err)
	} else {
		fmt.Printf("access granted with wrong password\n")
	}
	err, id = clientSUP.CheckUserPass("bob", "love2oscar")
	if err != nil {
		fmt.Printf("access denied with good password: %v\n", err)
	} else {
		fmt.Printf("access granted with ok password\n")
	}

	err = clientSUP.DeleteUser("bob")
	if err != nil {
		log.Fatalf("error deleting user: %v", err)
	}
	err, id = clientSUP.CheckUserPass("bob", "love2oscar")
	if err != nil {
		fmt.Printf("access denied for user password: %v\n", err)
	} else {
		fmt.Printf("access granted with deleted user\n")
	}
}
