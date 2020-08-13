package graphics

// Font declares all methods required to draw a font.
type Font interface {
	Close() error
	Draw(text string, x, y int32, color Color) error
}
