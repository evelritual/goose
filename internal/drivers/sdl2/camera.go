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

// Camera wraps methods used to alter the offsets of the current rendering
// context.
type Camera struct{}

// NewCamera sets the default offsets based on the current screen size.
func (s *SDL2) NewCamera() (graphics.Camera, error) {
	offsetX = windowX / 2
	offsetY = windowY / 2
	return &Camera{}, nil
}

// SetPosition sets the position of the camera relative to the world.
func (c *Camera) SetPosition(x, y int32) {
	offsetX = -x + (windowX / 2)
	offsetY = -y + (windowY / 2)
}

// SetScale sets the scale factor of all rendering textures. This is
// multiplicative of the actual texture scale size.
func (c *Camera) SetScale(x, y float32) {
	scaleFactorX = x
	scaleFactorY = y
}

// ScaleX returns the current X-axis scale factor of the camera.
func (c *Camera) ScaleX() float32 {
	return scaleFactorX
}

// ScaleY returns the current Y-axis scale factor of the camera.
func (c *Camera) ScaleY() float32 {
	return scaleFactorY
}

// X returns the current X offset in pixels relative to the world.
func (c *Camera) X() int32 {
	return offsetX
}

// Y returns the current Y offset in pixels relative to the world.
func (c *Camera) Y() int32 {
	return offsetY
}
