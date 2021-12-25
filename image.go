package main

import (
	"bytes"
	"image"
	"image/jpeg"
)

func convert(img *image.RGBA) (file *bytes.Buffer) {
	file = new(bytes.Buffer)
	_ = jpeg.Encode(file, img, nil)
	return
}
