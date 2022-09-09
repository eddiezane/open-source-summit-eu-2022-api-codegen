from concurrent import futures
from uuid import uuid4 as uuid
from threading import Lock

import grpc

from txt2img.v1 import txt2img_pb2
from txt2img.v1 import txt2img_pb2_grpc
from infer import run_inference
from storage import upload_image, delete_image
    

class Txt2ImgService(txt2img_pb2_grpc.Txt2ImgServiceServicer):
    def __init__(self):
        self.lock = Lock()
        self.images = {}

    def GenerateImage(self, request, context):
        prompt = request.prompt
        if len(prompt) == 0:
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "prompt is required")
        image = run_inference(prompt)
        id = str(uuid())
        url = upload_image(id, image)
        with self.lock:
            self.images[id] = url
        return txt2img_pb2.GenerateImageResponse(id=id)

    def GetImage(self, request, context):
        id = request.id
        if len(id) == 0:
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "id is required")
        with self.lock:
            if id not in self.images:
                context.abort(grpc.StatusCode.INVALID_ARGUMENT, "id not found")
            url = self.images[id]
        return txt2img_pb2.GetImageResponse(url=url)

    def ListImages(self, request, context):
        with self.lock:
            images = self.images.keys()
        return txt2img_pb2.ListImagesResponse(images=images)

    def DeleteImage(self, request, context):
        id = request.id
        if len(id) == 0:
            context.abort(grpc.StatusCode.INVALID_ARGUMENT, "id is required")
        with self.lock:
            if id not in self.images:
                context.abort(grpc.StatusCode.INVALID_ARGUMENT, "id not found")
            del self.images[id]
        delete_image(id)
        return txt2img_pb2.DeleteImageResponse()
            

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    txt2img_pb2_grpc.add_Txt2ImgServiceServicer_to_server(Txt2ImgService(), server)
    server.add_insecure_port("[::]:8080")
    server.start()
    server.wait_for_termination()

if __name__ == "__main__":
    print("starting server...")
    serve()
