# Copyright (c) 2018 Catena cyber
# Author Philippe Antoine <p.antoine@catenacyber.fr>
# Python test client for catenaSUP

import grpc

import userpass_pb2
import userpass_pb2_grpc

def run():
    channel = grpc.insecure_channel('localhost:1541')
    stub = userpass_pb2_grpc.CatenaUserPassStub(channel)
    oscarup = userpass_pb2.UserPass(user="oscar", password="sneakysnake")
    id = stub.AddUser(oscarup)
    print id


if __name__ == '__main__':
    run()
