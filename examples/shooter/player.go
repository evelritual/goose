package main

import (
	"github.com/evelritual/goose/graphics"
)

// Player ...
type Player struct {
	tex graphics.Texture

	speed int32
	x, y  int32
}
