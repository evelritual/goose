package main

import (
	"github.com/PapayaJuice/goose/internal/texture"
)

// Enemy ...
type Enemy struct {
	destroyed bool
	tex       texture.Texture

	speed int32
	x, y  int32
}
