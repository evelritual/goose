package goose

import "github.com/PapayaJuice/goose/internal/input"

// NewKeyboard initializes a new keyboard for use based on the active driver.
func NewKeyboard() input.Keyboard {
	return activeDriver.NewKeyboard()
}
