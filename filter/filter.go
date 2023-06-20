package filter

import (
	"bytes"
	"errors"
	"image"
	"image/color"

	"github.com/chai2010/webp"
)

func Filter(img image.Image, filter string) (bytes.Buffer, error) {
	// Handle the apply filter endpoint

	buf := bytes.Buffer{}

	var newImg image.Image

	switch filter {
	case "grayscale":
		var im *image.Gray
		im = image.NewGray(img.Bounds())
		for x := 0; x < img.Bounds().Max.X; x++ {
			for y := 0; y < img.Bounds().Max.Y; y++ {
				// r, g, b, _ := img.At(x, y).RGBA()
				// newImg.Set(x, y, color.Gray{uint8((0.2126 * float32(r)) + (0.2152 * float32(g)) + (0.0722 * float32(b)))})
				im.Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
			}
		}
		newImg = im
	case "sepia":
		var im *image.RGBA
		im = image.NewRGBA(img.Bounds())
		for x := 0; x < img.Bounds().Max.X; x++ {
			for y := 0; y < img.Bounds().Max.Y; y++ {
				r, g, b, _ := img.At(x, y).RGBA()
				// fmt.Println(r, g, b)
				newR := (0.393 * float64(r)) + (0.769 * float64(g)) + (0.189 * float64(b))
				newG := (0.349 * float64(r)) + (0.686 * float64(g)) + (0.168 * float64(b))
				newB := (0.272 * float64(r)) + (0.534 * float64(g)) + (0.131 * float64(b))
				im.Set(x, y, color.RGBA{uint8(newR / 256), uint8(newG / 256), uint8(newB / 256), 255})
			}
		}
		newImg = im
	default:
		return buf, errors.New("invalid filter")
	}

	webp.Encode(&buf, newImg, &webp.Options{Lossless: true})

	return buf, nil
}
