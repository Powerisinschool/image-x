# ImageX

This is an API for image conversion, built using Go and the Gin framework.

## API Documentation

The API documentation is generated using Swagger/OpenAPI and can be accessed at the following endpoint:
[localhost/docs/index.html](http://localhost:8080/docs/index.html)


## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Powerisinschool/image-x.git
   ```

2. Install the dependencies:

   ```bash
   go mod download
   ```

## Usage

1. Start the API server:

   ```bash
   go run main.go
   ```

2. Access the API endpoints using a REST client or tools like cURL or Postman.

## API Endpoints

- `POST /api/convert/format`: Convert image format.
- `POST /api/resize`: Resize image.
- `POST /api/filter`: Apply filter to image.
- `POST /api/crop`: Crop image.
- `POST /api/rotate`: Rotate image.
- `POST /api/compress`: Compress image.
- `POST /api/thumbnail`: Generate thumbnail image.

## Examples

### Convert Image Format

Convert the image format by sending a POST request to '/api/convert/format' endpoint with the following parameters:

- `file`: Image file to convert (multipart/form-data)
- `format`: Target format to convert to (multipart/form-data)

Example request using cURL:

```bash
curl -X POST -F "file=@/path/to/image.jpg" -F "format=png" http://localhost:8080/api/convert/format
```

### Resize Image

Resize the image by sending a POST request to '/api/resize' endpoint with the following parameters:

- `file`: Image file to resize (multipart/form-data)
- `width`: Target width for resizing (multipart/form-data)
- `height`: Target height for resizing (multipart/form-data)
- `quality`: Quality level for resizing. Options: `low`, `medium`, `high`, `best` (multipart/form-data)

Example request using cURL:

```bash
curl -X POST -F "file=@/path/to/image.jpg" -F "width=800" -F "height=600" -F "quality=80" http://localhost:8080/api/resize
```
