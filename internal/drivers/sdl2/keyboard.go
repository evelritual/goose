package sdl2

import (
	"github.com/PapayaJuice/goose/input"
)

// Keyboard wraps a map of SDL keys and their current states.
type Keyboard struct {
	keyStates map[input.Key]*input.KeyState
}

// NewKeyboard sets up an empty map of key states.
func (s *SDL2) NewKeyboard() input.Keyboard {
	s.keyboard = &Keyboard{
		keyStates: map[input.Key]*input.KeyState{},
	}
	return s.keyboard
}

// IsKeyDown returns true if the given key is pressed.
func (k *Keyboard) IsKeyDown(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return k.Pressed
	}
	return false
}

// IsKeyUp returns true of the given key is not pressed.
func (k *Keyboard) IsKeyUp(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return !k.Pressed
	}
	return false
}

// IsKeyPress returns true if the given key is pressed for the first time.
func (k *Keyboard) IsKeyPress(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return k.Pressed && !k.Repeat
	}
	return false
}

// IsKeyRelease returns true if the given key is pressed for the first time.
func (k *Keyboard) IsKeyRelease(keyCode input.Key) bool {
	if k, exists := k.keyStates[keyCode]; exists {
		return !k.Pressed && !k.Repeat
	}
	return false
}

// UpdateKey updates the state of the given key to the keyboard state map.
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
