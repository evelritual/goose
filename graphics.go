package goose

import (
	"github.com/PapayaJuice/goose/graphics"
)

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

// NewTextureAtlus loads a new texture and splices it into drawable tiles.
// Texture must be closed after use.
func NewTextureAtlus(imgPath string, splitX, splitY int32) (graphics.TextureAtlus, error) {
	return activeDriver.NewTextureAtlus(imgPath, splitX, splitY)
}
