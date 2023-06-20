package converter

import (
	"bytes"
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/Powerisinschool/image-x/match"
	"github.com/chai2010/webp"
)

var supportedFormats = []string{"png", "jpeg", "webp", "gif"}

// Pixel struct example
type Pixel struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func Convert(img image.Image, format string) (bytes.Buffer, error) {

	buf := bytes.Buffer{}

	if !match.IsMatching(format, supportedFormats) {
		return buf, errors.New("Unsupported format: image/" + format)
	}

	switch format {
	case "png":
		png.Encode(&buf, img)
	case "jpeg":
		jpeg.Encode(&buf, img, &jpeg.Options{Quality: 100})
	case "webp":
		webp.Encode(&buf, img, &webp.Options{Lossless: true})
	case "gif":
		gif.Encode(&buf, img, nil)
	}

	return buf, nil
}
