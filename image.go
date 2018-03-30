package main

import (
	"image"
	"image/png"
	"os"
)

// Image - I wonder what you are!
type Image struct {
	image *image.RGBA
}

// NewImage creates a new image of the give size.
func NewImage(width, height int) *Image {
	return &Image{
		image: image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

// WriteToFile saves an image to a file.
func (i *Image) WriteToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	if err := png.Encode(f, i.image); err != nil {
		return err
	}

	f.Close()
	return nil
}
