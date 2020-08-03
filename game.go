package goose

const (
	defaultImage = "goose.png"
)

// Game declares all methods required to run a game
type Game interface {
	Draw()
	Update() error
}

type defaultGame struct {
	tex *Texture
}

// Draw ...
func (d *defaultGame) Draw() {}

// Update ...
func (d *defaultGame) Update() error {
	return nil
}
