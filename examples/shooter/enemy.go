package main

import (
	"github.com/PapayaJuice/goose/graphics"
)

// Enemy ...
type Enemy struct {
	destroyed bool
	tex       graphics.Texture

	speed int32
	x, y  int32
}
