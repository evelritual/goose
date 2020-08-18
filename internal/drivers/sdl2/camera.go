package sdl2

import (
	"github.com/PapayaJuice/goose/graphics"
)

var (
	offsetX      int32   = 0
	offsetY      int32   = 0
	scaleFactorX float32 = 1.0
	scaleFactorY float32 = 1.0
)

// Camera ...
type Camera struct{}

// NewCamera ...
func (s *SDL2) NewCamera() (graphics.Camera, error) {
	offsetX = windowX / 2
	offsetY = windowY / 2
	return &Camera{}, nil
}

// SetPosition ...
func (c *Camera) SetPosition(x, y int32) {
	offsetX = -x + (windowX / 2)
	offsetY = -y + (windowY / 2)
}

// SetScale ...
func (c *Camera) SetScale(x, y float32) {
	scaleFactorX = x
	scaleFactorY = y
}

// ScaleX ...
func (c *Camera) ScaleX() float32 {
	return scaleFactorX
}

// ScaleY ...
func (c *Camera) ScaleY() float32 {
	return scaleFactorY
}

// X ...
func (c *Camera) X() int32 {
	return offsetX
}

// Y ...
func (c *Camera) Y() int32 {
	return offsetY
}
