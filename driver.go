package goose

import (
	"log"

	"github.com/PapayaJuice/goose/internal/input"

	"github.com/PapayaJuice/goose/drivers/sdl2"
	"github.com/PapayaJuice/goose/internal/texture"
)

const (
	// DriverSDL2 represents the SDL2 library
	DriverSDL2 = "sdl2"
)

var (
	activeDriver driver
	driverMap    = map[string]driver{
		DriverSDL2: &sdl2.SDL2{},
	}
)

// driver declares all methods needed to implement a functional driver. Drivers
// are used for basic interactions between the system and Goose such as drawing
// to screen, input, audio, and more.
type driver interface {
	Init() error
	CreateWindow(x, y int32, title string) error
	Close()
	NewKeyboard() input.Keyboard
	NewTexture(imgPath string) (texture.Texture, error)
	PreDraw() error
	PostDraw()
	Update() error
}

// setDriver sets the given driver as the default.
func setDriver(driver string) {
	d, exists := driverMap[driver]
	if !exists {
		log.Fatalf("no such driver: %s", driver)
	}

	err := d.Init()
	if err != nil {
		log.Fatalf("error initializing driver %s: %v", driver, err)
	}

	activeDriver = d
}
