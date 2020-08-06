package goose

import (
	"github.com/PapayaJuice/goose/internal/texture"
)

// NewTexture loads a new texture for use in drawing. Texture must be closed
// after use.
func NewTexture(imgPath string) (texture.Texture, error) {
	return activeDriver.NewTexture(imgPath)
}
