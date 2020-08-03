package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Texture ...
type Texture struct {
	image   *sdl.Surface
	texture *sdl.Texture

	x int32
	y int32
}

// NewTexture ...
func (s *SDL2) NewTexture(image string) (*Texture, error) {
	img, err := img.Load(image)
	if err != nil {
		return nil, fmt.Errorf("error loading image: %v", err)
	}
	b := img.Bounds()

	tex, err := s.renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("error creating texture: %v", err)
	}

	return &Texture{
		image:   img,
		texture: tex,
		x:       int32(b.Size().X),
		y:       int32(b.Size().Y),
	}, nil
}

// Close ...
func (t *Texture) Close() {
	t.texture.Destroy()
	t.image.Free()
}
