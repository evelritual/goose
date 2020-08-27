package input

// MouseButton represents a single button on a mouse.
type MouseButton int

// Mouse declares all methods required to interact with a mouse.
type Mouse interface {
	IsButtonDown(button MouseButton) bool
	IsButtonPress(button MouseButton) bool
	IsButtonRelease(button MouseButton) bool
	IsButtonUp(button MouseButton) bool
	IsClick(button MouseButton) (bool, int32, int32)
	IsRelease(button MouseButton) (bool, int32, int32)
	UpdateButton(button MouseButton, x, y int32, pressed bool)
	UpdatePosition(x, y int32)
	Pos() (int32, int32)
}
