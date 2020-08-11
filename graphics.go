package goose

import (
	"github.com/PapayaJuice/goose/graphics"
)

// NewTexture loads a new texture for use in drawing. Texture must be closed
// after use.
func NewTexture(imgPath string) (graphics.Texture, error) {
	return activeDriver.NewTexture(imgPath)
}
