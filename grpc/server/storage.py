from google.cloud import storage

client = storage.Client()
bucket = client.bucket("eddiezane-chainguard-pub")

def upload_image(id, image):
    blob = bucket.blob(f"{id}.png")
    blob.upload_from_file(image, content_type="image/png")
    return blob.public_url

def delete_image(id):
    blob = bucket.blob(f"{id}.png")
    blob.delete()
