package main

import (
	"path"
)

const imagePath = "./testimages/star-wars.jpg"
const outputDir = "resizedimages"

func main() {
	resized, err := resizeImage(imagePath)
	if err != nil {
		println("Failed to resize image", err)
		return
	}

	if err := writeImageToFile(resized, path.Join(outputDir, "resized.jpg")); err != nil {
		println("Failed saving resized image", err)
	}
}
