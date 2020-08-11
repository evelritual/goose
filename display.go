package goose

var (
	windowTitle = "Goose Engine"
	windowX     = int32(320)
	windowY     = int32(240)
)

// SetWindowSize sets the x and y of the screen before creation. It has no
// effect after a window is created.
func SetWindowSize(x, y int32) {
	windowX = x
	windowY = y
}

// GetWindowX returns the current width of the window.
func GetWindowX() int32 {
	return windowX
}

// GetWindowY returns the current height of the window.
func GetWindowY() int32 {
	return windowY
}

// SetWindowTitle sets the title of the window before creation. It has no
// effect after a window is created.
func SetWindowTitle(title string) {
	windowTitle = title
}
