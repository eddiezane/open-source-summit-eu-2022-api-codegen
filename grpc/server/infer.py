import io
import torch
from torch import autocast
from diffusers import StableDiffusionPipeline

assert torch.cuda.is_available()

pipe = StableDiffusionPipeline.from_pretrained(
        'CompVis/stable-diffusion-v1-4',
        use_auth_token=True
).to('cuda')

def run_inference(prompt):
  with autocast('cuda'):
      image = pipe(prompt)['sample'][0]
  img_data = io.BytesIO()
  image.save(img_data, 'PNG')
  img_data.seek(0)
  return img_data
