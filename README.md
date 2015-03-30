# spritesheet

Command-line utility to generate sprite sheet based on individual sprite images written in Go.

This utility can be used to generate sprite sheet textures to be used in games. The output images will be constrainted to a specified width and height (2048x2048 by default). This is useful when working with programs that output sprites as individual images (Blender, for example).

![example sprite sheet](http://i.imgur.com/hjukpGF.png)

# Usage

    spritesheet [indir] [outdir]

# Installation

This utility is written in [Go](http://www.golang.org/), so you need to have that installed first. Once go is install, simply run:
    
    go get github.com/jamal/spritesheet

And provided your go/bin directory is in your path, you should just be able to run `spritesheet`.
