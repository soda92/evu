package images

import (
	"image"

	"github.com/gohugoio/hugo/resources/images/exif"
	"github.com/gohugoio/hugo/resources/resource"
)

// ImageResource represents an image resource.
type ImageResource interface {
	resource.Resource
	ImageResourceOps
}

type ImageResourceOps interface {
	// Height returns the height of the Image.
	Height() int

	// Width returns the width of the Image.
	Width() int

	// Process applies the given image processing options to the image.
	Process(spec string) (ImageResource, error)

	// Crop an image to match the given dimensions without resizing.
	// You must provide both width and height.
	// Use the anchor option to change the crop box anchor point.
	//    {{ $image := $image.Crop "600x400" }}
	Crop(spec string) (ImageResource, error)

	// Fill scales the image to the smallest possible size that will cover the specified dimensions in spec,
	// crops the resized image to the specified dimensions using the given anchor point.
	// The spec is space delimited, e.g. `200x300 TopLeft`.
	Fill(spec string) (ImageResource, error)

	// Fit scales down the image using the given spec.
	Fit(spec string) (ImageResource, error)

	// Resize resizes the image to the given spec. If one of width or height is 0, the image aspect
	// ratio is preserved.
	Resize(spec string) (ImageResource, error)

	// Filter applies one or more filters to an Image.
	//    {{ $image := $image.Filter (images.GaussianBlur 6) (images.Pixelate 8) }}
	Filter(filters ...any) (ImageResource, error)

	// Exif returns an ExifInfo object containing Image metadata.
	Exif() *exif.ExifInfo

	// Colors returns a slice of the most dominant colors in an image
	// using a simple histogram method.
	Colors() ([]Color, error)

	// For internal use.
	DecodeImage() (image.Image, error)
}
