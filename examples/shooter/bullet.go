package main

import (
	"github.com/PapayaJuice/goose/graphics"
)

// Bullet ...
type Bullet struct {
	destroyed bool
	tex       graphics.Texture

	speed int32
	x, y  int32
}
