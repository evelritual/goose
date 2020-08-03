package goose

import (
	"fmt"

	"github.com/PapayaJuice/goose/drivers/sdl2"
)

const (
	// DriverSDL2 represents the SDL2 library
	DriverSDL2 = "sdl2"
)

var (
	driverMap = map[string]Driver{
		DriverSDL2: &sdl2.SDL2{},
	}
)

// Driver declares all methods needed to implement a functional driver. Drivers
// are used for basic interactions between the system and Goose such as drawing
//  to screen, input, audio, and more.
type Driver interface {
	Init() error
	CreateWindow(x, y int, title string) error
	Close()
	PreDraw() error
	PostDraw()
	Update() error
}

func getDriver(driver string) (Driver, error) {
	d, exists := driverMap[driver]
	if !exists {
		return nil, fmt.Errorf("no such driver: %s", driver)
	}

	err := d.Init()
	if err != nil {
		return nil, fmt.Errorf("error initializing driver %s: %v", driver, err)
	}
	return d, nil
}
