package loadx

import (
	"image"
	"mime/multipart"
)

func Load(file multipart.File) (image.Image, error) {
	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
