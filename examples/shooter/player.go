package main

import (
	"github.com/PapayaJuice/goose/internal/texture"
)

// Player ...
type Player struct {
	tex texture.Texture

	speed int32
	x, y  int32
}
