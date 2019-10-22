package synthesis

import "errors"

// Generator contains information required for generation
type Generator struct {
	srcFile string
	dstFile string
}

// New creates a new generator
func New(src, dst string) *Generator {
	return &Generator{
		srcFile: src,
		dstFile: dst,
	}
}

// Load loads the source image from disk
func (g *Generator) Load() error {
	return errors.New("Not implemented")
}

// Generate perfoms the generation with the supplied random seed
func (g *Generator) Generate(seed int) error {
	return errors.New("Not implemented")
}

// Save saves the generated image to disk
func (g *Generator) Save() error {
	return errors.New("Not implemented")
}
