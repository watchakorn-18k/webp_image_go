## Image to WebP Converter (Golang)

A lightweight, efficient backend service built in Golang to convert images to WebP format. This project includes a Dockerfile for easy deployment and scalability.

## Build Docker

```bash
# Podman
podman build -t image_api -f Dockerfile .

# Docker
docker build -t image_api -f Dockerfile .
```

## API Endpoint

URL: `/convert-to-webp`
Method: `POST`
Content-Type: `multipart/form-data`
Description: Uploads an image file and returns the converted WebP file.
Body : image (file)

CURL:

```bash
curl --location 'http://127.0.0.1:3000/convert-to-webp' \
--form 'image=@"/path/to/file"'
```
