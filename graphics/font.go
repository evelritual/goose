package graphics

// Font declares all methods required to draw a font.
type Font interface {
	SetFont(fontPath string, size int) error
	Texture(text string, color Color) (Texture, error)
	Close() error
}
