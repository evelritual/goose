package graphics

var (
	// ColorBlack is the color black.
	ColorBlack = Color{R: 0, G: 0, B: 0, A: 0}
	// ColorBlue is the color blue.
	ColorBlue = Color{R: 0, G: 0, B: 255, A: 0}
	// ColorGreen is the color green.
	ColorGreen = Color{R: 0, G: 255, B: 0, A: 0}
	// ColorRed is the color red.
	ColorRed = Color{R: 255, G: 0, B: 0, A: 0}
	// ColorWhite is the color white.
	ColorWhite = Color{R: 255, G: 255, B: 255, A: 0}
)

// Color represents a single RGBA color.
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}
