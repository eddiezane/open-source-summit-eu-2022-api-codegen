# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: txt2img/v1/txt2img.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18txt2img/v1/txt2img.proto\x12\ntxt2img.v1\x1a\x1cgoogle/api/annotations.proto\"A\n\x05Image\x12\x16\n\x06prompt\x18\x01 \x01(\tR\x06prompt\x12\x0e\n\x02id\x18\x02 \x01(\tR\x02id\x12\x10\n\x03url\x18\x03 \x01(\tR\x03url\".\n\x14GenerateImageRequest\x12\x16\n\x06prompt\x18\x01 \x01(\tR\x06prompt\"@\n\x15GenerateImageResponse\x12\'\n\x05image\x18\x01 \x01(\x0b\x32\x11.txt2img.v1.ImageR\x05image\"!\n\x0fGetImageRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\";\n\x10GetImageResponse\x12\'\n\x05image\x18\x01 \x01(\x0b\x32\x11.txt2img.v1.ImageR\x05image\"\x13\n\x11ListImagesRequest\"?\n\x12ListImagesResponse\x12)\n\x06images\x18\x01 \x03(\x0b\x32\x11.txt2img.v1.ImageR\x06images\"$\n\x12\x44\x65leteImageRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"\x15\n\x13\x44\x65leteImageResponse2\xa7\x03\n\x0eTxt2ImgService\x12k\n\rGenerateImage\x12 .txt2img.v1.GenerateImageRequest\x1a!.txt2img.v1.GenerateImageResponse\"\x15\x82\xd3\xe4\x93\x02\x0f:\x01*\"\n/v1/images\x12^\n\x08GetImage\x12\x1b.txt2img.v1.GetImageRequest\x1a\x1c.txt2img.v1.GetImageResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/v1/images/{id}\x12_\n\nListImages\x12\x1d.txt2img.v1.ListImagesRequest\x1a\x1e.txt2img.v1.ListImagesResponse\"\x12\x82\xd3\xe4\x93\x02\x0c\x12\n/v1/images\x12g\n\x0b\x44\x65leteImage\x12\x1e.txt2img.v1.DeleteImageRequest\x1a\x1f.txt2img.v1.DeleteImageResponse\"\x17\x82\xd3\xe4\x93\x02\x11*\x0f/v1/images/{id}B^Z\\github.com/eddiezane/open-source-summit-eu-2022-api-codegen/grpc/client/txt2img/v1;txt2imgv1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'txt2img.v1.txt2img_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\\github.com/eddiezane/open-source-summit-eu-2022-api-codegen/grpc/client/txt2img/v1;txt2imgv1'
  _TXT2IMGSERVICE.methods_by_name['GenerateImage']._options = None
  _TXT2IMGSERVICE.methods_by_name['GenerateImage']._serialized_options = b'\202\323\344\223\002\017:\001*\"\n/v1/images'
  _TXT2IMGSERVICE.methods_by_name['GetImage']._options = None
  _TXT2IMGSERVICE.methods_by_name['GetImage']._serialized_options = b'\202\323\344\223\002\021\022\017/v1/images/{id}'
  _TXT2IMGSERVICE.methods_by_name['ListImages']._options = None
  _TXT2IMGSERVICE.methods_by_name['ListImages']._serialized_options = b'\202\323\344\223\002\014\022\n/v1/images'
  _TXT2IMGSERVICE.methods_by_name['DeleteImage']._options = None
  _TXT2IMGSERVICE.methods_by_name['DeleteImage']._serialized_options = b'\202\323\344\223\002\021*\017/v1/images/{id}'
  _IMAGE._serialized_start=70
  _IMAGE._serialized_end=135
  _GENERATEIMAGEREQUEST._serialized_start=137
  _GENERATEIMAGEREQUEST._serialized_end=183
  _GENERATEIMAGERESPONSE._serialized_start=185
  _GENERATEIMAGERESPONSE._serialized_end=249
  _GETIMAGEREQUEST._serialized_start=251
  _GETIMAGEREQUEST._serialized_end=284
  _GETIMAGERESPONSE._serialized_start=286
  _GETIMAGERESPONSE._serialized_end=345
  _LISTIMAGESREQUEST._serialized_start=347
  _LISTIMAGESREQUEST._serialized_end=366
  _LISTIMAGESRESPONSE._serialized_start=368
  _LISTIMAGESRESPONSE._serialized_end=431
  _DELETEIMAGEREQUEST._serialized_start=433
  _DELETEIMAGEREQUEST._serialized_end=469
  _DELETEIMAGERESPONSE._serialized_start=471
  _DELETEIMAGERESPONSE._serialized_end=492
  _TXT2IMGSERVICE._serialized_start=495
  _TXT2IMGSERVICE._serialized_end=918
# @@protoc_insertion_point(module_scope)
