package input

const (
	// MouseButtonLeft "Left Mouse Button"
	MouseButtonLeft = MouseButton(0)
	// MouseButtonMiddle "Middle Mouse Button"
	MouseButtonMiddle = MouseButton(1)
	// MouseButtonRight "Right Mouse Button"
	MouseButtonRight = MouseButton(2)
)

// MouseButtonState represents the current state of a mouse button.
type MouseButtonState struct {
	Pressed bool
	Repeat  bool
	X       int32
	Y       int32
}
