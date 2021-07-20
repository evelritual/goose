package main

import (
	"github.com/evelritual/goose/graphics"
)

// Enemy ...
type Enemy struct {
	destroyed bool
	tex       graphics.Texture

	speed int32
	x, y  int32
}
