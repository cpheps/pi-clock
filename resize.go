package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

func resizeImage(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	originalImg, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	resizedImg := resize.Resize(64, 0, originalImg, resize.Lanczos3)

	return resizedImg, nil
}

func writeImageToFile(imageData image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, imageData, nil)
}
