package goose

import (
	"log"

	"github.com/PapayaJuice/goose/audio"
	"github.com/PapayaJuice/goose/graphics"
	"github.com/PapayaJuice/goose/input"
	"github.com/PapayaJuice/goose/internal/drivers/sdl2"
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
	NewAudioPlayer() (audio.Player, error)
	NewCamera() (graphics.Camera, error)
	NewFont(fontPath string, size int) (graphics.Font, error)
	NewKeyboard() input.Keyboard
	NewTexture(imgPath string) (graphics.Texture, error)
	NewTextureAtlus(imgPath string, splitX, splitY int32) (graphics.TextureAtlus, error)
	PreDraw() error
	PostDraw()
	SetBackgroundColor(color *graphics.Color)
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
