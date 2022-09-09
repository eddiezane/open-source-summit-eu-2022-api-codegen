import grpc
from txt2img.v1 import txt2img_pb2
from txt2img.v1 import txt2img_pb2_grpc


channel = grpc.insecure_channel("34.135.240.67:8080")
client = txt2img_pb2_grpc.Txt2ImgServiceStub(channel)
res = client.ListImages(txt2img_pb2.ListImagesRequest())
print(f"{res}")

for id in res.images:
    res = client.DeleteImage(txt2img_pb2.DeleteImageRequest(id=id))
    print(f"{res}")
