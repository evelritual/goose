package goose

import (
	"fmt"

	"github.com/evelritual/goose/graphics"
	"github.com/evelritual/goose/input"
)

const (
	defaultImage = "../../logo.png"
)

// Game declares all methods required to run a game
type Game interface {
	Close() error
	Draw() error
	FixedUpdate() error
	Init() error
	Update() error
}

type defaultGame struct {
	keyboard input.Keyboard

	tex        graphics.Texture
	shouldDraw bool
	texX       int32
	texY       int32
	speedX     int32
	speedY     int32
}

// Init ...
func (d *defaultGame) Init() error {
	SetBackgroundColor(&graphics.ColorWhite)
	d.keyboard = NewKeyboard()
	t, err := NewTexture(defaultImage)
	if err != nil {
		return fmt.Errorf("error loading default image: %v", err)
	}
	d.tex = t
	d.shouldDraw = true

	d.texX = (windowX / 2) - (d.tex.W() / 16)
	d.texY = (windowY / 2) - (d.tex.H() / 16)

	d.speedX = 3
	d.speedY = 2

	return nil
}

// Draw ...
func (d *defaultGame) Draw() error {
	if d.tex == nil {
		return nil
	}

	if !d.shouldDraw {
		return nil
	}

	err := d.tex.Draw(d.texX, d.texY, 0.125, 0.125, 0.0)
	if err != nil {
		return fmt.Errorf("error drawing default image: %v", err)
	}
	return nil
}

// FixedUpdate ..
func (d *defaultGame) FixedUpdate() error {
	if d.texX <= 0 || d.texX+(d.tex.W()/8) >= windowX {
		d.speedX = 0 - d.speedX
	}
	if d.texY <= 0 || d.texY+(d.tex.H()/8) >= windowY {
		d.speedY = 0 - d.speedY
	}

	d.texX += d.speedX
	d.texY += d.speedY

	return nil
}

// Update ...
func (d *defaultGame) Update() error {
	if d.keyboard.IsKeyPress(input.KeySpace) {
		d.shouldDraw = false
	}
	if d.keyboard.IsKeyRelease(input.KeySpace) {
		d.shouldDraw = true
	}

	return nil
}

// Close ...
func (d *defaultGame) Close() error {
	err := d.tex.Close()
	if err != nil {
		return fmt.Errorf("error closing game: %v", err)
	}
	return nil
}
