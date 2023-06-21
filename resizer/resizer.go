package resizer

import (
	"bytes"
	"errors"
	"image"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

func Resize(img image.Image, width, height int, quality string) (bytes.Buffer, error) {

	buf := bytes.Buffer{}

	// Set the expected width and height of the image
	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	var scaler draw.Interpolator

	switch quality {
	case "low":
		scaler = draw.NearestNeighbor
	case "medium":
		scaler = draw.ApproxBiLinear
	case "high":
		scaler = draw.BiLinear
	case "best":
		scaler = draw.CatmullRom
	default:
		return buf, errors.New("invalid quality")
	}

	// Resize the image to the dst image's size and return it
	scaler.Scale(dst, dst.Rect, img, img.Bounds(), draw.Over, nil)

	// Encode the resized image to webp
	webp.Encode(&buf, dst, &webp.Options{Lossless: true})

	return buf, nil
}
