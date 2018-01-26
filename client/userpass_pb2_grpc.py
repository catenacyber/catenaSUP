# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import userpass_pb2 as userpass__pb2


class CatenaUserPassStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.AddUser = channel.unary_unary(
        '/catenaSUP.CatenaUserPass/AddUser',
        request_serializer=userpass__pb2.UserPass.SerializeToString,
        response_deserializer=userpass__pb2.Id.FromString,
        )
    self.ChangePass = channel.unary_unary(
        '/catenaSUP.CatenaUserPass/ChangePass',
        request_serializer=userpass__pb2.UserPass.SerializeToString,
        response_deserializer=userpass__pb2.Empty.FromString,
        )
    self.CheckUserPass = channel.unary_unary(
        '/catenaSUP.CatenaUserPass/CheckUserPass',
        request_serializer=userpass__pb2.UserPass.SerializeToString,
        response_deserializer=userpass__pb2.Id.FromString,
        )
    self.DeleteUser = channel.unary_unary(
        '/catenaSUP.CatenaUserPass/DeleteUser',
        request_serializer=userpass__pb2.User.SerializeToString,
        response_deserializer=userpass__pb2.Empty.FromString,
        )


class CatenaUserPassServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def AddUser(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ChangePass(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CheckUserPass(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteUser(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_CatenaUserPassServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'AddUser': grpc.unary_unary_rpc_method_handler(
          servicer.AddUser,
          request_deserializer=userpass__pb2.UserPass.FromString,
          response_serializer=userpass__pb2.Id.SerializeToString,
      ),
      'ChangePass': grpc.unary_unary_rpc_method_handler(
          servicer.ChangePass,
          request_deserializer=userpass__pb2.UserPass.FromString,
          response_serializer=userpass__pb2.Empty.SerializeToString,
      ),
      'CheckUserPass': grpc.unary_unary_rpc_method_handler(
          servicer.CheckUserPass,
          request_deserializer=userpass__pb2.UserPass.FromString,
          response_serializer=userpass__pb2.Id.SerializeToString,
      ),
      'DeleteUser': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteUser,
          request_deserializer=userpass__pb2.User.FromString,
          response_serializer=userpass__pb2.Empty.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'catenaSUP.CatenaUserPass', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))