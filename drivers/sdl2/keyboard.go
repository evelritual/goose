package sdl2

import (
	"github.com/PapayaJuice/goose/internal/input"
)

// Keyboard ...
type Keyboard struct {
	keyStates map[input.Key]*input.KeyState
}

// NewKeyboard ...
func (s *SDL2) NewKeyboard() input.Keyboard {
	s.keyboard = &Keyboard{
		keyStates: map[input.Key]*input.KeyState{},
	}
	return s.keyboard
}

// IsKeyDown ...
func (k *Keyboard) IsKeyDown(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return k.Pressed
	}
	return false
}

// IsKeyUp ...
func (k *Keyboard) IsKeyUp(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return !k.Pressed
	}
	return false
}

// IsKeyPress ...
func (k *Keyboard) IsKeyPress(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return k.Pressed && !k.Repeat
	}
	return false
}

// IsKeyRelease ...
func (k *Keyboard) IsKeyRelease(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return !k.Pressed && !k.Repeat
	}
	return false
}

// UpdateKey ...
func (k *Keyboard) UpdateKey(keyCode input.Key, pressed, repeat bool) {
	if k, exists := k.keyStates[keyCode]; exists {
		k.Pressed = pressed
		k.Repeat = repeat
		return
	}

	k.keyStates[keyCode] = &input.KeyState{
		Pressed: pressed,
		Repeat:  repeat,
	}
}
