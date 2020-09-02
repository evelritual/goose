package sdl2

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// DisableCursor disables rendering of the hardware cursor.
func (s *SDL2) DisableCursor() error {
	_, err := sdl.ShowCursor(sdl.DISABLE)
	if err != nil {
		return fmt.Errorf("error setting sdl show cursor to disabled: %v", err)
	}
	return nil
}

// EnableCursor enables rendering of the hardware cursor.
func (s *SDL2) EnableCursor() error {
	_, err := sdl.ShowCursor(sdl.ENABLE)
	if err != nil {
		return fmt.Errorf("error setting sdl show cursor to enabled: %v", err)
	}
	return nil
}
