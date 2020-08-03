package goose

import (
	"fmt"
	"runtime"
	"time"
)

var (
	targetDriver = DriverSDL2
	targetFPS    = 60
	windowX      = 320
	windowY      = 240
	windowTitle  = "Goose Engine"
)

// Run starts the main game loop of the Goose engine using the Update and Draw
// methods provided by the given Game object.
func Run(game Game) error {
	runtime.LockOSThread()
	if game == nil {
		game = &defaultGame{}
	}

	// Init window
	d, err := getDriver(targetDriver)
	if err != nil {
		return fmt.Errorf("error getting driver: %v", err)
	}
	defer d.Close()

	err = d.CreateWindow(windowX, windowY, windowTitle)
	if err != nil {
		return fmt.Errorf("error creating window: %v", err)
	}

	// Run loop
	// Cap at target FPS
	// TODO Allow when targetFPS = 0 (unlimited)
	fpst := 1000 / targetFPS

	for range time.Tick(time.Duration(fpst) * time.Millisecond) {
		// Update driver state first
		err = d.Update()
		if err != nil {
			break
		}

		err = game.Update()
		if err != nil {
			break
		}

		err = d.PreDraw()
		if err != nil {
			break
		}
		game.Draw()
		d.PostDraw()
	}
	return err
}
