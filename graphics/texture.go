package graphics

// Texture declares all methods required to draw a texture.
type Texture interface {
	Close() error
	Draw(x, y int32, scaleX, scaleY float32, rotation float64) error
	W() int32
	H() int32
}

// TextureAtlas declares all methods required to draw a texture atlas.
type TextureAtlas interface {
	Close() error
	Draw(tile int, x, y int32, scaleX, scaleY float32, rotation float64) error
	Len() int
}
