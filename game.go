package goose

import (
	"fmt"

	"github.com/PapayaJuice/goose/texture"
)

const (
	defaultImage = "../../goose.png"
)

// Game declares all methods required to run a game
type Game interface {
	Draw() error
	Update() error
}

type defaultGame struct {
	tex texture.Texture
}

// Draw ...
func (d *defaultGame) Draw() error {
	if d.tex == nil {
		return nil
	}

	x := (windowX / 2) - (d.tex.W() / 8)
	y := (windowY / 2) - (d.tex.H() / 8)
	err := d.tex.Draw(x, y, 0.25, 0.25)
	if err != nil {
		return fmt.Errorf("error drawing default image: %v", err)
	}
	return nil
}

// Update ...
func (d *defaultGame) Update() error {
	if d.tex == nil {
		t, err := NewTexture(defaultImage)
		if err != nil {
			return fmt.Errorf("error loading default image: %v", err)
		}
		d.tex = t
	}

	return nil
}
