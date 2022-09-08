import grpc
from txt2img.v1 import txt2img_pb2
from txt2img.v1 import txt2img_pb2_grpc


channel = grpc.insecure_channel('34.135.240.67:8080')
client = txt2img_pb2_grpc.Txt2ImgServiceStub(channel)
req = txt2img_pb2.GenerateImageRequest(prompt='a photograph of a gopher wearing a wizard hat in a forest, vivid, photorealistic, magical, fantasy, 8K UHD, photography')
print(f'{req}')
res = client.GenerateImage(req)
print(f'{res}')

req = txt2img_pb2.GenerateImageRequest(prompt='an orange aligator')
print(f'{req}')
res = client.GenerateImage(req)

res = client.ListImages(txt2img_pb2.ListImagesRequest())
print(f'{res}')
id = res.images[0]

res = client.GetImage(txt2img_pb2.GetImageRequest(id=id))
print(f'{res}')

res = client.DeleteImage(txt2img_pb2.DeleteImageRequest(id=id))
print(f'{res}')