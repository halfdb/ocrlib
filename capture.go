package main

import (
	"github.com/kbinani/screenshot"
	"image"
)

func shot(x, y, w, h int) *image.RGBA {
	rect, err := screenshot.CaptureRect(image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x + w,
			Y: y + h,
		},
	})
	if err != nil {
		return nil
	}
	return rect
}
