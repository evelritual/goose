package texture

// Texture declares all methods required to draw a texture.
type Texture interface {
	Close() error
	Draw(x, y int32, scaleX, scaleY float32) error
	W() int32
	H() int32
}
