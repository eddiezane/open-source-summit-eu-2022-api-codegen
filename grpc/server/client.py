import os
import grpc
from txt2img.v1 import txt2img_pb2
from txt2img.v1 import txt2img_pb2_grpc

host = os.getenv("TXT2IMG_HOST")
if host == "" or host == None:
    raise Exception("TXT2IMG_HOST required")

channel = grpc.insecure_channel(host)
client = txt2img_pb2_grpc.Txt2ImgServiceStub(channel)
req = txt2img_pb2.GenerateImageRequest(prompt="an orange cat")
print(f"{req}")
res = client.GenerateImage(req)
print(f"{res}")

req = txt2img_pb2.GenerateImageRequest(prompt="an orange aligator")
print(f"{req}")
res = client.GenerateImage(req)

res = client.ListImages(txt2img_pb2.ListImagesRequest())
print(f"{res}")
id = res.images[0]

res = client.GetImage(txt2img_pb2.GetImageRequest(id=id))
print(f"{res}")

res = client.DeleteImage(txt2img_pb2.DeleteImageRequest(id=id))
print(f"{res}")
