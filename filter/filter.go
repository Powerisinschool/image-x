package filter

import (
	"bytes"
	"errors"
	"image"
	"image/color"

	"github.com/chai2010/webp"
)

type FilterOptions struct {
	Radius int `json:"radius"`
}

func Filter(img image.Image, filter string, options *FilterOptions) (bytes.Buffer, error) {
	// Handle the apply filter endpoint

	buf := bytes.Buffer{}

	var newImg image.Image

	switch filter {
	case "grayscale":
		var im *image.Gray
		im = image.NewGray(img.Bounds())
		for x := 0; x < img.Bounds().Max.X; x++ {
			for y := 0; y < img.Bounds().Max.Y; y++ {
				r, g, b, _ := img.At(x, y).RGBA()
				im.Set(x, y, color.Gray{uint8(((0.299 * float32(r)) + (0.587 * float32(g)) + (0.114 * float32(b))) / 256)})
				// im.Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
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

				// Divide RGB values by factor to remove noise from the result
				var factor float64 = 352 // (256 + 64 + 32)
				im.Set(x, y, color.RGBA{uint8(newR / factor), uint8(newG / factor), uint8(newB / factor), 255})
			}
		}
		newImg = im
	case "blur":
		var im *image.RGBA
		im = image.NewRGBA(img.Bounds())
		for x := 0; x < img.Bounds().Max.X; x++ {
			for y := 0; y < img.Bounds().Max.Y; y++ {
				r, g, b, _ := img.At(x, y).RGBA()
				im.Set(x, y, color.RGBA{uint8(r / 256), uint8(g / 256), uint8(b / 256), 255})
			}
		}
		newImg = applyBokehBlur(im, options.Radius)
	default:
		return buf, errors.New("invalid filter")
	}

	webp.Encode(&buf, newImg, &webp.Options{Lossless: true})

	return buf, nil
}

func applyBokehBlur(imageRGB *image.RGBA, blurRadius int) *image.RGBA {
	resultImage := image.NewRGBA(imageRGB.Bounds())

	for x := imageRGB.Bounds().Min.X; x < imageRGB.Bounds().Max.X; x++ {
		for y := imageRGB.Bounds().Min.Y; y < imageRGB.Bounds().Max.Y; y++ {
			avgColor := calculateAverageColor(imageRGB, x, y, blurRadius)
			resultImage.SetRGBA(x, y, avgColor)
		}
	}

	return resultImage
}

func calculateAverageColor(imageRGB *image.RGBA, x, y, blurRadius int) color.RGBA {
	var totalR, totalG, totalB uint32
	numPixels := 0

	for i := x - blurRadius; i <= x+blurRadius; i++ {
		for j := y - blurRadius; j <= y+blurRadius; j++ {
			if isWithinBounds(imageRGB, i, j) {
				r, g, b, _ := imageRGB.At(i, j).RGBA()
				totalR += r
				totalG += g
				totalB += b
				numPixels++
			}
		}
	}

	avgR := uint8(totalR / uint32(numPixels) >> 8)
	avgG := uint8(totalG / uint32(numPixels) >> 8)
	avgB := uint8(totalB / uint32(numPixels) >> 8)

	return color.RGBA{avgR, avgG, avgB, 255}
}

func isWithinBounds(imageRGB *image.RGBA, x, y int) bool {
	bounds := imageRGB.Bounds()
	return x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y
}

// func applyBokehFilter(img image.Image, blurRadius int) *image.RGBA {
// 	var im *image.RGBA
// 	im = image.NewRGBA(img.Bounds())
// 	for x := 0; x < img.Bounds().Max.X; x++ {
// 		for y := 0; y < img.Bounds().Max.Y; y++ {
// 			// r, g, b, _ = img.At(x, y).RGBA()
// 			avgColor := getAverageColor(img, x, y, blurRadius)
// 			im.Set(x, y, avgColor)
// 		}
// 	}
// 	return im
// }

// func getAverageColor(img image.Image, x int, y int, blurRadius int) color.RGBA {
// 	var r, g, b, a uint32
// 	var count int
// 	for i := x - blurRadius; i < x+blurRadius; i++ {
// 		for j := y - blurRadius; j < y+blurRadius; j++ {
// 			if i < 0 || j < 0 || i >= img.Bounds().Max.X || j >= img.Bounds().Max.Y {
// 				continue
// 			}
// 			r, g, b, a = img.At(i, j).RGBA()
// 			count++
// 		}
// 	}
// 	var factor uint32 = 64
// 	return color.RGBA{uint8((r / uint32(count)) / factor), uint8((g / uint32(count)) / factor), uint8((b / uint32(count)) / factor), uint8(a / uint32(count))}
// }
