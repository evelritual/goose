package goose

import (
	"github.com/evelritual/goose/graphics"
)

// NewCamera creates a new camera for use in rendering the viewport in an
// altered state. Initial position starts at 0,0.
func NewCamera() (graphics.Camera, error) {
	return activeDriver.NewCamera()
}

// NewFont loads a new ttf font for use in drawing text. Font must be closed
// after use.
func NewFont(fontPath string, size int) (graphics.Font, error) {
	return activeDriver.NewFont(fontPath, size)
}

// NewTexture loads a new texture for use in drawing. Texture must be closed
// after use.
func NewTexture(imgPath string) (graphics.Texture, error) {
	return activeDriver.NewTexture(imgPath)
}

// NewTextureAtlas loads a new texture and splices it into drawable tiles.
// Texture must be closed after use.
func NewTextureAtlas(imgPath string, splitX, splitY int32) (graphics.TextureAtlas, error) {
	return activeDriver.NewTextureAtlas(imgPath, splitX, splitY)
}
