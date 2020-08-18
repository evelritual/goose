package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/PapayaJuice/goose/graphics"
)

// TextureAtlus wraps an SDL Texture. It splits the texture into tiles to
// allow for easy drawing of sprites in a spritesheet.
type TextureAtlus struct {
	renderer *sdl.Renderer // reference to renderer to use

	image   *sdl.Surface
	texture *sdl.Texture

	tiles []*sdl.Rect
	tileH int32
	tileW int32
}

// NewTextureAtlus loads an image as an SDL texture and splices it into
// separate rectangles for use in drawing.
func (s *SDL2) NewTextureAtlus(imgPath string, splitX, splitY int32) (graphics.TextureAtlus, error) {
	img, err := img.Load(imgPath)
	if err != nil {
		return nil, fmt.Errorf("error loading image: %v", err)
	}

	// Split up image
	// Doesn't count for remainder
	b := img.Bounds().Size()
	w := int32(b.X) / splitX
	h := int32(b.Y) / splitY

	tiles := []*sdl.Rect{}
	for y := int32(0); y < h; y++ {
		for x := int32(0); x < w; x++ {
			tiles = append(tiles, &sdl.Rect{X: x * splitX, Y: y * splitY, W: splitX, H: splitY})
		}
	}

	tex, err := s.renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("error creating texture: %v", err)
	}

	return &TextureAtlus{
		renderer: s.renderer,
		image:    img,
		texture:  tex,
		tiles:    tiles,
		tileH:    splitY,
		tileW:    splitX,
	}, nil
}

// Draw renders the texture of the given tile to the SDL renderer.
func (t *TextureAtlus) Draw(tile int, x, y int32, scaleX, scaleY float32) error {
	if tile > len(t.tiles)-1 {
		return fmt.Errorf("tile out of range")
	}

	err := t.renderer.Copy(
		t.texture,
		t.tiles[tile],
		&sdl.Rect{
			X: x,
			Y: y,
			W: int32(float32(t.tileW) * scaleX),
			H: int32(float32(t.tileH) * scaleY),
		},
	)

	if err != nil {
		return fmt.Errorf("error drawing: %v", err)
	}
	return nil
}

// Len returns the number of spliced tiles in the texture atlus.
func (t *TextureAtlus) Len() int {
	return len(t.tiles)
}

// Close releases the texture and the image from memory.
func (t *TextureAtlus) Close() error {
	err := t.texture.Destroy()
	if err != nil {
		return fmt.Errorf("error destroying sdl texture: %v", err)
	}

	t.image.Free()
	return nil
}
