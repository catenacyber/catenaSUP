# CatenaSUP

A simple tool to secure your users passwords for your web application

## Install and setup

CatenaSUP works with a RPC server and client.
The server should be run on antoher machine, or at least with another user account, to ensure data can only be accessed through the RPC.

Setup the communication :
- a port where to bind the server
- one certificate for client and one for the server

To setup the server, you will need a database file.
You can init one with :
`go run initdb.go file.db`

Then you can run the server
`go run server.go`
Its options are :
- certificate file for client
- certificate and key file for server
- database file
- port to listen

To setup the client, you need to add/modify your code where the passwords chacking takes place, to call the RPC.

An example is available and you can test it with :
`go run dummyCli.go`

All the client has to do is :
- open the connection to the server
- call when needed the RPC procedures AddUser, ChangePass, CheckUserPass, DeleteUser
- close the connection

## License

GPLv3

This software uses :
- grpc (Apache License 2.0)
- protobuf (MIT-like license)
- sqlite3 (public domain)
- memguard (Apache License 2.0)
- argon2 (MIT License)

## TODO

- Automate testing
- Clients in other languages
- Easy installation, import existing table...
