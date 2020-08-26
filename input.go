package goose

import "github.com/PapayaJuice/goose/input"

// NewKeyboard initializes a new keyboard for use based on the active driver.
func NewKeyboard() input.Keyboard {
	return activeDriver.NewKeyboard()
}

// NewMouse initializes a new mouse for use based on the active driver.
func NewMouse() input.Mouse {
	return activeDriver.NewMouse()
}
