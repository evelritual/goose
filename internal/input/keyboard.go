package input

// Key represents a single key on a keyboard.
type Key int

// Keyboard declares all methods required to interact with a keyboard.
type Keyboard interface {
	IsKeyDown(keyCode Key) bool
	IsKeyPress(keyCode Key) bool
	IsKeyRelease(keyCode Key) bool
	IsKeyUp(keyCode Key) bool
	UpdateKey(keyCode Key, pressed, repeat bool)
}
