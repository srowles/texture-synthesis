package synthesis

import (
	"image"
	"image/jpeg"
	"log"
	"os"
)

// Generator contains information required for generation
type Generator struct {
	example image.Image
	output  image.Image
}

// New creates a new generator
func New() *Generator {
	return &Generator{}
}

// Load loads the source image from disk
func (g *Generator) Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		log.Printf("Failed to open example file (%s) with error: %v", filename, err)
		return err
	}
	g.example, err = jpeg.Decode(f)
	if err != nil {
		log.Printf("Failed to decode jpg (%s) with error: %v", filename, err)
		return err
	}
	return nil
}

// Generate perfoms the generation with the supplied random seed
func (g *Generator) Generate(seed int) error {
	p := buildPyramid(g.example)
	g.output = p.images[2]
	return nil
}

// Save saves the generated image to disk
func (g *Generator) Save(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Printf("Failed to open output file (%s) with error: %v", filename, err)
		return err
	}

	err = jpeg.Encode(f, g.output, nil)
	if err != nil {
		log.Printf("Failed to encode jpg (%s) with error: %v", filename, err)
	}
	return nil
}
