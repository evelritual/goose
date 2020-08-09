package main

import (
	"github.com/PapayaJuice/goose/internal/texture"
)

// Bullet ...
type Bullet struct {
	destroyed bool
	tex       texture.Texture

	speed int32
	x, y  int32
}
