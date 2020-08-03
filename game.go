package goose

// Game declares all methods required to run a game
type Game interface {
	Draw()
	Update() error
}

type defaultGame struct{}

// Draw ...
func (d *defaultGame) Draw() {}

// Update ...
func (d *defaultGame) Update() error {
	return nil
}
