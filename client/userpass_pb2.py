# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: userpass.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='userpass.proto',
  package='catenaSUP',
  syntax='proto3',
  serialized_pb=_b('\n\x0euserpass.proto\x12\tcatenaSUP\"*\n\x08UserPass\x12\x0c\n\x04user\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x14\n\x04User\x12\x0c\n\x04user\x18\x01 \x01(\t\"\x07\n\x05\x45mpty\"\x10\n\x02Id\x12\n\n\x02id\x18\x01 \x01(\x04\x32\xe2\x01\n\x0e\x43\x61tenaUserPass\x12/\n\x07\x41\x64\x64User\x12\x13.catenaSUP.UserPass\x1a\r.catenaSUP.Id\"\x00\x12\x35\n\nChangePass\x12\x13.catenaSUP.UserPass\x1a\x10.catenaSUP.Empty\"\x00\x12\x35\n\rCheckUserPass\x12\x13.catenaSUP.UserPass\x1a\r.catenaSUP.Id\"\x00\x12\x31\n\nDeleteUser\x12\x0f.catenaSUP.User\x1a\x10.catenaSUP.Empty\"\x00\x62\x06proto3')
)




_USERPASS = _descriptor.Descriptor(
  name='UserPass',
  full_name='catenaSUP.UserPass',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='catenaSUP.UserPass.user', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='password', full_name='catenaSUP.UserPass.password', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=29,
  serialized_end=71,
)


_USER = _descriptor.Descriptor(
  name='User',
  full_name='catenaSUP.User',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='user', full_name='catenaSUP.User.user', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=73,
  serialized_end=93,
)


_EMPTY = _descriptor.Descriptor(
  name='Empty',
  full_name='catenaSUP.Empty',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=95,
  serialized_end=102,
)


_ID = _descriptor.Descriptor(
  name='Id',
  full_name='catenaSUP.Id',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='catenaSUP.Id.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=104,
  serialized_end=120,
)

DESCRIPTOR.message_types_by_name['UserPass'] = _USERPASS
DESCRIPTOR.message_types_by_name['User'] = _USER
DESCRIPTOR.message_types_by_name['Empty'] = _EMPTY
DESCRIPTOR.message_types_by_name['Id'] = _ID
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

UserPass = _reflection.GeneratedProtocolMessageType('UserPass', (_message.Message,), dict(
  DESCRIPTOR = _USERPASS,
  __module__ = 'userpass_pb2'
  # @@protoc_insertion_point(class_scope:catenaSUP.UserPass)
  ))
_sym_db.RegisterMessage(UserPass)

User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), dict(
  DESCRIPTOR = _USER,
  __module__ = 'userpass_pb2'
  # @@protoc_insertion_point(class_scope:catenaSUP.User)
  ))
_sym_db.RegisterMessage(User)

Empty = _reflection.GeneratedProtocolMessageType('Empty', (_message.Message,), dict(
  DESCRIPTOR = _EMPTY,
  __module__ = 'userpass_pb2'
  # @@protoc_insertion_point(class_scope:catenaSUP.Empty)
  ))
_sym_db.RegisterMessage(Empty)

Id = _reflection.GeneratedProtocolMessageType('Id', (_message.Message,), dict(
  DESCRIPTOR = _ID,
  __module__ = 'userpass_pb2'
  # @@protoc_insertion_point(class_scope:catenaSUP.Id)
  ))
_sym_db.RegisterMessage(Id)



_CATENAUSERPASS = _descriptor.ServiceDescriptor(
  name='CatenaUserPass',
  full_name='catenaSUP.CatenaUserPass',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=123,
  serialized_end=349,
  methods=[
  _descriptor.MethodDescriptor(
    name='AddUser',
    full_name='catenaSUP.CatenaUserPass.AddUser',
    index=0,
    containing_service=None,
    input_type=_USERPASS,
    output_type=_ID,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ChangePass',
    full_name='catenaSUP.CatenaUserPass.ChangePass',
    index=1,
    containing_service=None,
    input_type=_USERPASS,
    output_type=_EMPTY,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CheckUserPass',
    full_name='catenaSUP.CatenaUserPass.CheckUserPass',
    index=2,
    containing_service=None,
    input_type=_USERPASS,
    output_type=_ID,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteUser',
    full_name='catenaSUP.CatenaUserPass.DeleteUser',
    index=3,
    containing_service=None,
    input_type=_USER,
    output_type=_EMPTY,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CATENAUSERPASS)

DESCRIPTOR.services_by_name['CatenaUserPass'] = _CATENAUSERPASS

# @@protoc_insertion_point(module_scope)
