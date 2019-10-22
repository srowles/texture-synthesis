package main

import (
	"flag"
	"log"
	"os"

	synthesis "github.com/srowles/texture-synthesis"
)

func main() {
	var in, out string
	var seed int
	flag.StringVar(&in, "in", "", "Path to input file")
	flag.StringVar(&out, "out", "", "Outfile file")
	flag.IntVar(&seed, "seed", 0, "The seed to use for image generation")
	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			flag.Usage()
			os.Exit(1)
		}
	})

	generator := synthesis.New(in, out)
	if err := generator.Load(); err != nil {
		log.Fatalf("Failed to load example file from disk: %v", err)
	}

	if err := generator.Generate(seed); err != nil {
		log.Fatalf("Failed to generate image with error: %v", err)
	}

	if err := generator.Save(); err != nil {
		log.Fatalf("Failed to save generated image to disk: %v", err)
	}
}
