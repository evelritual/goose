package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// SDL2 implements the Driver interface
type SDL2 struct {
	renderer *sdl.Renderer
	window   *sdl.Window

	keyboard *Keyboard
}

// Init initializes everything in the SDL2 library
func (s *SDL2) Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("error initializing sdl: %v", err)
	}
	return nil
}

// CreateWindow starts a new window and renderer which is ready to draw to
func (s *SDL2) CreateWindow(x, y int32, title string) error {
	w, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, x, y, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("error creating sdl window: %v", err)
	}
	s.window = w

	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("error creating sdl renderer: %v", err)
	}
	s.renderer = r

	err = s.PreDraw()
	if err != nil {
		return fmt.Errorf("error with predraw: %v", err)
	}
	s.PostDraw()

	return nil
}

// PreDraw flushes the renderer and sets up for drawing
func (s *SDL2) PreDraw() error {
	err := s.renderer.SetDrawColor(0, 0, 0, 0)
	if err != nil {
		return fmt.Errorf("error setting renderer draw color: %v", err)
	}

	err = s.renderer.Clear()
	if err != nil {
		return fmt.Errorf("error clearing renderer: %v", err)
	}

	return nil
}

// PostDraw writes all new bytes to the renderer
func (s *SDL2) PostDraw() {
	s.renderer.Present()
}

// Update parses through all new events SDL reports on and updates the state
// accordingly.
func (s *SDL2) Update() error {
	// Reset all keys that were just pressed
	for _, s := range s.keyboard.keyStates {
		s.Repeat = true
	}

	// Poll SDL events
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return fmt.Errorf("exit")
		case *sdl.KeyboardEvent:
			if s.keyboard == nil {
				break
			}
			// TODO handle unknown key press
			s.keyboard.UpdateKey(keyMap[t.Keysym.Sym], t.State == sdl.PRESSED, t.Repeat != 0)
		}
	}
	return nil
}

// Close safely releases SDL resources.
func (s *SDL2) Close() {
	s.renderer.Destroy()
	s.window.Destroy()
}