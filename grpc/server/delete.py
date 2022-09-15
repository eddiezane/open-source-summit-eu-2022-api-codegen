import os
import grpc
from txt2img.v1 import txt2img_pb2
from txt2img.v1 import txt2img_pb2_grpc

host = os.getenv("TXT2IMG_HOST")
if host == "" or host == None:
    raise Exception("TXT2IMG_HOST required")

channel = grpc.insecure_channel(host)
client = txt2img_pb2_grpc.Txt2ImgServiceStub(channel)
res = client.ListImages(txt2img_pb2.ListImagesRequest())
print(f"{res}")

for img in res.images:
    res = client.DeleteImage(txt2img_pb2.DeleteImageRequest(id=img.id))
    print(f"{res}")
