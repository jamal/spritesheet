package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"path"
	"path/filepath"
)

var inDir string
var outDir string
var width int
var height int

func init() {
	flag.IntVar(&width, "width", 2048, "width of the output image(s)")
	flag.IntVar(&height, "height", 2048, "height of the output image(s)")

	if len(os.Args) < 3 {
		usage()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [indir] [outdir]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func writeSheet(outDir string, c int, im image.Image) error {
	out, err := os.Create(path.Join(outDir, fmt.Sprintf("%d.png", c)))
	if err != nil {
		return err
	}

	err = png.Encode(out, im)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	inDir := os.Args[1]
	outDir := os.Args[2]

	files, err := filepath.Glob(path.Join(inDir, "*.*"))
	if err != nil {
		panic(err)
	}

	c := 0

	var xoff, yoff int
	im := image.NewRGBA(image.Rect(0, 0, width, height))

	for _, filename := range files {
		r, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file %s: %s\n", filename, err.Error())
			os.Exit(2)
		}

		cim, _, err := image.Decode(r)
		// Ignore files that are not images, but fail otherwise
		if err != nil && err != image.ErrFormat {
			fmt.Fprintf(os.Stderr, "failed to decode image %s: %s\n", filename, err.Error())
			os.Exit(2)
		}

		if err != image.ErrFormat {
			dp := image.Pt(xoff, yoff)
			draw.Draw(im, image.Rectangle{dp, dp.Add(cim.Bounds().Size())}, cim, cim.Bounds().Min, draw.Src)

			xoff = xoff + cim.Bounds().Size().X
			if xoff >= width {
				xoff = 0
				yoff = yoff + cim.Bounds().Size().Y
				if yoff >= height {
					yoff = 0

					err = writeSheet(outDir, c, im)
					if err != nil {
						fmt.Fprintf(os.Stderr, "failed to write sprite sheet: %s\n", err.Error())
						os.Exit(2)
					}

					c = c + 1
					im = image.NewRGBA(image.Rect(0, 0, width, height))
				}
			}
		}
	}

	err = writeSheet(outDir, c, im)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to write sprite sheet: %s\n", err.Error())
		os.Exit(2)
	}
}
