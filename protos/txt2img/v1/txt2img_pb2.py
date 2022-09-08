# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: txt2img/v1/txt2img.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18txt2img/v1/txt2img.proto\x12\ntxt2img.v1\"&\n\x14GenerateImageRequest\x12\x0e\n\x06prompt\x18\x01 \x01(\t\"#\n\x15GenerateImageResponse\x12\n\n\x02id\x18\x01 \x01(\t\"\x1d\n\x0fGetImageRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x1f\n\x10GetImageResponse\x12\x0b\n\x03url\x18\x01 \x01(\t\"\x13\n\x11ListImagesRequest\"$\n\x12ListImagesResponse\x12\x0e\n\x06images\x18\x01 \x03(\t\" \n\x12\x44\x65leteImageRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x15\n\x13\x44\x65leteImageResponse2\xd2\x02\n\x0eTxt2ImgService\x12V\n\rGenerateImage\x12 .txt2img.v1.GenerateImageRequest\x1a!.txt2img.v1.GenerateImageResponse\"\x00\x12G\n\x08GetImage\x12\x1b.txt2img.v1.GetImageRequest\x1a\x1c.txt2img.v1.GetImageResponse\"\x00\x12M\n\nListImages\x12\x1d.txt2img.v1.ListImagesRequest\x1a\x1e.txt2img.v1.ListImagesResponse\"\x00\x12P\n\x0b\x44\x65leteImage\x12\x1e.txt2img.v1.DeleteImageRequest\x1a\x1f.txt2img.v1.DeleteImageResponse\"\x00\x42OZMgithub.com/eddiezane/open-source-summit-eu-2022-api-codegen/protos/txt2img/v1b\x06proto3')



_GENERATEIMAGEREQUEST = DESCRIPTOR.message_types_by_name['GenerateImageRequest']
_GENERATEIMAGERESPONSE = DESCRIPTOR.message_types_by_name['GenerateImageResponse']
_GETIMAGEREQUEST = DESCRIPTOR.message_types_by_name['GetImageRequest']
_GETIMAGERESPONSE = DESCRIPTOR.message_types_by_name['GetImageResponse']
_LISTIMAGESREQUEST = DESCRIPTOR.message_types_by_name['ListImagesRequest']
_LISTIMAGESRESPONSE = DESCRIPTOR.message_types_by_name['ListImagesResponse']
_DELETEIMAGEREQUEST = DESCRIPTOR.message_types_by_name['DeleteImageRequest']
_DELETEIMAGERESPONSE = DESCRIPTOR.message_types_by_name['DeleteImageResponse']
GenerateImageRequest = _reflection.GeneratedProtocolMessageType('GenerateImageRequest', (_message.Message,), {
  'DESCRIPTOR' : _GENERATEIMAGEREQUEST,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.GenerateImageRequest)
  })
_sym_db.RegisterMessage(GenerateImageRequest)

GenerateImageResponse = _reflection.GeneratedProtocolMessageType('GenerateImageResponse', (_message.Message,), {
  'DESCRIPTOR' : _GENERATEIMAGERESPONSE,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.GenerateImageResponse)
  })
_sym_db.RegisterMessage(GenerateImageResponse)

GetImageRequest = _reflection.GeneratedProtocolMessageType('GetImageRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETIMAGEREQUEST,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.GetImageRequest)
  })
_sym_db.RegisterMessage(GetImageRequest)

GetImageResponse = _reflection.GeneratedProtocolMessageType('GetImageResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETIMAGERESPONSE,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.GetImageResponse)
  })
_sym_db.RegisterMessage(GetImageResponse)

ListImagesRequest = _reflection.GeneratedProtocolMessageType('ListImagesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTIMAGESREQUEST,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.ListImagesRequest)
  })
_sym_db.RegisterMessage(ListImagesRequest)

ListImagesResponse = _reflection.GeneratedProtocolMessageType('ListImagesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTIMAGESRESPONSE,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.ListImagesResponse)
  })
_sym_db.RegisterMessage(ListImagesResponse)

DeleteImageRequest = _reflection.GeneratedProtocolMessageType('DeleteImageRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEIMAGEREQUEST,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.DeleteImageRequest)
  })
_sym_db.RegisterMessage(DeleteImageRequest)

DeleteImageResponse = _reflection.GeneratedProtocolMessageType('DeleteImageResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETEIMAGERESPONSE,
  '__module__' : 'txt2img.v1.txt2img_pb2'
  # @@protoc_insertion_point(class_scope:txt2img.v1.DeleteImageResponse)
  })
_sym_db.RegisterMessage(DeleteImageResponse)

_TXT2IMGSERVICE = DESCRIPTOR.services_by_name['Txt2ImgService']
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'ZMgithub.com/eddiezane/open-source-summit-eu-2022-api-codegen/protos/txt2img/v1'
  _GENERATEIMAGEREQUEST._serialized_start=40
  _GENERATEIMAGEREQUEST._serialized_end=78
  _GENERATEIMAGERESPONSE._serialized_start=80
  _GENERATEIMAGERESPONSE._serialized_end=115
  _GETIMAGEREQUEST._serialized_start=117
  _GETIMAGEREQUEST._serialized_end=146
  _GETIMAGERESPONSE._serialized_start=148
  _GETIMAGERESPONSE._serialized_end=179
  _LISTIMAGESREQUEST._serialized_start=181
  _LISTIMAGESREQUEST._serialized_end=200
  _LISTIMAGESRESPONSE._serialized_start=202
  _LISTIMAGESRESPONSE._serialized_end=238
  _DELETEIMAGEREQUEST._serialized_start=240
  _DELETEIMAGEREQUEST._serialized_end=272
  _DELETEIMAGERESPONSE._serialized_start=274
  _DELETEIMAGERESPONSE._serialized_end=295
  _TXT2IMGSERVICE._serialized_start=298
  _TXT2IMGSERVICE._serialized_end=636
# @@protoc_insertion_point(module_scope)