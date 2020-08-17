package graphics

// Texture declares all methods required to draw a texture.
type Texture interface {
	Close() error
	Draw(x, y int32, scaleX, scaleY float32) error
	W() int32
	H() int32
}

// TextureAtlus declares all methods required to draw a texture atlus.
type TextureAtlus interface {
	Close() error
	Draw(tile int, x, y int32, scaleX, scaleY float32) error
	Len() int
}
