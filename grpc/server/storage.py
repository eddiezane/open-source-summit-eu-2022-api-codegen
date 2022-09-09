import os
from google.cloud import storage

bucket_name = os.getenv("TXT2IMG_BUCKET")

if bucket_name == "" or bucket_name == None:
    raise Exception("TXT2IMG_BUCKET required")

client = storage.Client()
bucket = client.bucket(bucket_name)

def upload_image(id, image):
    blob = bucket.blob(f"{id}.png")
    blob.upload_from_file(image, content_type="image/png")
    return blob.public_url

def delete_image(id):
    blob = bucket.blob(f"{id}.png")
    blob.delete()
