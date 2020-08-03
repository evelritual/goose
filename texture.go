package goose

// Texture ...
type Texture interface {
	Close()
	Draw()
}

// LoadTexture ...
func LoadTexture(path string) (Texture, error) {
	return nil, nil
}
