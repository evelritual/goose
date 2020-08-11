package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/PapayaJuice/goose/graphics"
)

// Texture ...
type Texture struct {
	renderer *sdl.Renderer // reference to renderer to use

	image   *sdl.Surface
	texture *sdl.Texture

	w int32
	h int32
}

// NewTexture loads an image as an SDL texture. Texture must be closed
// manually.
func (s *SDL2) NewTexture(imgPath string) (graphics.Texture, error) {
	img, err := img.Load(imgPath)
	if err != nil {
		return nil, fmt.Errorf("error loading image: %v", err)
	}
	b := img.Bounds()

	tex, err := s.renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("error creating texture: %v", err)
	}

	return &Texture{
		renderer: s.renderer,
		image:    img,
		texture:  tex,
		w:        int32(b.Size().X),
		h:        int32(b.Size().Y),
	}, nil
}

// W returns the width of the texture.
func (t *Texture) W() int32 {
	return t.w
}

// H returns the height of the texture.
func (t *Texture) H() int32 {
	return t.h
}

// Draw renders the texture to the SDL renderer.
func (t *Texture) Draw(x, y int32, scaleX, scaleY float32) error {
	err := t.renderer.Copy(t.texture,
		&sdl.Rect{
			X: 0,
			Y: 0,
			W: t.w,
			H: t.h,
		},
		&sdl.Rect{
			X: x,
			Y: y,
			W: int32(float32(t.w) * scaleX),
			H: int32(float32(t.h) * scaleY),
		},
	)

	if err != nil {
		return fmt.Errorf("error drawing: %v", err)
	}
	return nil
}

// Close releases the texture and the image from memory.
func (t *Texture) Close() error {
	err := t.texture.Destroy()
	if err != nil {
		return fmt.Errorf("error destroying sdl texture: %v", err)
	}

	t.image.Free()
	return nil
}
