package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"

	"github.com/veandco/go-sdl2/ttf"

	"github.com/PapayaJuice/goose/graphics"
)

// Font wraps an SDL TTF Font and allows drawing to screen.
type Font struct {
	renderer *sdl.Renderer // reference to renderer to use

	font *ttf.Font
}

// NewFont ...
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

// Draw ...
func (f *Font) Draw(text string, x, y int32, color graphics.Color) error {
	c := sdl.Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}

	t, err := f.font.RenderUTF8Blended(text, c)
	if err != nil {
		return fmt.Errorf("error rendering sdl text: %v", err)
	}
	defer t.Free()

	tex, err := f.renderer.CreateTextureFromSurface(t)
	if err != nil {
		return fmt.Errorf("error creating sdl surface for font: %v", err)
	}
	defer tex.Destroy()

	_, _, w, h, err := tex.Query()
	if err != nil {
		return fmt.Errorf("error querying sdl font info: %v", err)
	}

	err = f.renderer.Copy(tex, nil, &sdl.Rect{X: x, Y: y, W: w, H: h})
	if err != nil {
		return fmt.Errorf("error rendering sdl font on screen: %v", err)
	}
	return nil
}

// Close ...
func (f *Font) Close() error {
	f.font.Close()
	return nil
}
