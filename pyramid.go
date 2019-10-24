package synthesis

import (
	"image"
	"math"

	"github.com/esimov/stackblur-go"
	"github.com/nfnt/resize"
)

type pyramid struct {
	images []image.Image
	levels int
}

func buildPyramid(input image.Image) *pyramid {
	max := math.Max(float64(input.Bounds().Size().X), float64(input.Bounds().Size().Y))
	levels := int(math.Log2(max))
	p := pyramid{
		levels: levels,
		images: make([]image.Image, levels),
	}

	// level 0 = original image
	p.images[0] = input

	for level := 1; level < levels; level++ {
		input = gaussianDownSample(input)
		p.images[level] = input
	}

	return &p
}

func gaussianDownSample(input image.Image) image.Image {
	blured := stackblur.Process(input, 10)
	size := input.Bounds().Size()
	targetX := uint(size.X / 2)
	targetY := uint(size.Y / 2)
	return resize.Resize(targetX, targetY, blured, resize.Bilinear)
}
