package sdl2

import (
	"github.com/PapayaJuice/goose/input"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	mouseButtonMap = map[uint8]input.MouseButton{
		sdl.BUTTON_LEFT:   input.MouseButtonLeft,
		sdl.BUTTON_MIDDLE: input.MouseButtonMiddle,
		sdl.BUTTON_RIGHT:  input.MouseButtonRight,
	}
)

// Mouse wraps a map of SDL Mouse Buttons and their current states as well as
// the current mouse position.
type Mouse struct {
	buttonStates map[input.MouseButton]*input.MouseButtonState
	X            int32
	Y            int32
}

// NewMouse sets up an empty map of mouse button states.
func (s *SDL2) NewMouse() input.Mouse {
	s.mouse = &Mouse{
		buttonStates: map[input.MouseButton]*input.MouseButtonState{},
	}
	return s.mouse
}

// IsButtonDown returns true if the given mouse button is pressed.
func (m *Mouse) IsButtonDown(button input.MouseButton) bool {
	if b, exists := m.buttonStates[button]; exists {
		return b.Pressed
	}
	return false
}

// IsButtonUp returns true if the given mouse button is not pressed.
func (m *Mouse) IsButtonUp(button input.MouseButton) bool {
	if b, exists := m.buttonStates[button]; exists {
		return !b.Pressed
	}
	return false
}

// IsButtonPress returns true if the given mouse button is pressed for the
// first time.
func (m *Mouse) IsButtonPress(button input.MouseButton) bool {
	if b, exists := m.buttonStates[button]; exists {
		return b.Pressed && !b.Repeat
	}
	return false
}

// IsButtonRelease returns true if the given mouse button is released for the
// first time.
func (m *Mouse) IsButtonRelease(button input.MouseButton) bool {
	if b, exists := m.buttonStates[button]; exists {
		return !b.Pressed && !b.Repeat
	}
	return false
}

// IsClick returns true if the given mouse button is pressed for the first
// time as well as its clicked coordinates
func (m *Mouse) IsClick(button input.MouseButton) (bool, int32, int32) {
	if b, exists := m.buttonStates[button]; exists {
		return b.Pressed && !b.Repeat, b.X, b.Y
	}
	return false, 0, 0
}

// IsRelease returns true if the given mouse button is released for the first
// time as well as its released coordinates
func (m *Mouse) IsRelease(button input.MouseButton) (bool, int32, int32) {
	if b, exists := m.buttonStates[button]; exists {
		return !b.Pressed && !b.Repeat, b.X, b.Y
	}
	return false, 0, 0
}

// UpdateButton updates the state of the given button to the mouse button state
// map.
func (m *Mouse) UpdateButton(button input.MouseButton, x, y int32, pressed bool) {
	if b, exists := m.buttonStates[button]; exists {
		b.Pressed = pressed
		b.Repeat = false
		b.X = x
		b.Y = y
		return
	}

	m.buttonStates[button] = &input.MouseButtonState{
		Pressed: pressed,
		X:       x,
		Y:       y,
	}
}

// UpdatePosition updates the mouse cursor position.
func (m *Mouse) UpdatePosition(x, y int32) {
	m.X = x
	m.Y = y
}

// Pos returns the current X and Y position of the mouse.
func (m *Mouse) Pos() (int32, int32) {
	return m.X, m.Y
}
