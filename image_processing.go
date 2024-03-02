package imageprocessing

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// ReadImage now returns an image and an error.
func ReadImage(path string) (image.Image, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	// Decode the image
	img, _, err := image.Decode(inputFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// WriteImage now returns an error.
func WriteImage(path string, img image.Image) error {
	outputFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Encode the image to the new file
	err = jpeg.Encode(outputFile, img, nil)
	if err != nil {
		return err
	}
	return nil
}

// Grayscale now returns an image and an error.
func Grayscale(img image.Image) (image.Image, error) {
	// Create a new grayscale image
	bounds := img.Bounds()
	grayImg := image.NewGray(bounds)

	// Convert each pixel to grayscale
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalPixel := img.At(x, y)
			grayPixel := color.GrayModel.Convert(originalPixel)
			grayImg.Set(x, y, grayPixel)
		}
	}
	return grayImg, nil
}

// Resize now returns an image and an error.
func Resize(img image.Image) (image.Image, error) {
	newWidth := uint(500)
	newHeight := uint(500)
	resizedImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	if resizedImg == nil {
		return nil, fmt.Errorf("resize failed")
	}
	return resizedImg, nil
}
