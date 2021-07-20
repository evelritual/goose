package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	"github.com/evelritual/goose/graphics"
)

// Font wraps an SDL TTF Font and allows drawing to screen.
type Font struct {
	renderer *sdl.Renderer // reference to renderer to use

	font *ttf.Font
}

// NewFont opens a TTF font and sets it up for use in rendering.
func (s *SDL2) NewFont(fontPath string, size int) (graphics.Font, error) {
	f, err := ttf.OpenFont(fontPath, size)
	if err != nil {
		return nil, fmt.Errorf("error loading sdl font: %v", err)
	}

	return &Font{
		renderer: s.renderer,
		font:     f,
	}, nil
}

// SetFont loads a new font. All future calls to Texture will use the newly
// loaded font.
func (f *Font) SetFont(fontPath string, size int) error {
	font, err := ttf.OpenFont(fontPath, size)
	if err != nil {
		return fmt.Errorf("error loading sdl font: %v", err)
	}

	f2 := f.font
	defer f2.Close()

	f.font = font
	return nil
}

// Texture loads the font as a drawable texture. Texture must be closed
// manually.
func (f *Font) Texture(text string, color graphics.Color) (graphics.Texture, error) {
	c := sdl.Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}

	t, err := f.font.RenderUTF8Blended(text, c)
	if err != nil {
		return nil, fmt.Errorf("error rendering sdl text: %v", err)
	}
	defer t.Free()

	tex, err := f.renderer.CreateTextureFromSurface(t)
	if err != nil {
		return nil, fmt.Errorf("error creating sdl surface for font: %v", err)
	}

	_, _, w, h, err := tex.Query()
	if err != nil {
		return nil, fmt.Errorf("error querying sdl font info: %v", err)
	}

	return &Texture{
		renderer: f.renderer,
		image:    nil,
		texture:  tex,
		w:        w,
		h:        h,
	}, nil
}

// Close releases the SDL font resource.
func (f *Font) Close() error {
	f.font.Close()
	return nil
}
