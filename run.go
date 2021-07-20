package goose

import (
	"fmt"
	"runtime"
	"time"
)

var (
	targetDriver = DriverSDL2
	targetFPS    = 60
)

// Run starts the main game loop of the Goose engine using the Update and Draw
// methods provided by the given Game object.
func Run(game Game) error {
	runtime.LockOSThread()
	if game == nil {
		game = &defaultGame{}
	}

	// Init window
	setDriver(targetDriver)
	defer activeDriver.Close()

	err := activeDriver.CreateWindow(windowX, windowY, windowTitle)
	if err != nil {
		return fmt.Errorf("error creating window: %v", err)
	}

	// Init game
	err = game.Init()
	if err != nil {
		return fmt.Errorf("error initializing game: %v", err)
	}
	defer game.Close()

	// Run loop
	// Cap at target FPS
	// TODO Allow when targetFPS = 0 (unlimited)
	fpst := 1000 / targetFPS
	drawChan := time.Tick(time.Duration(fpst) * time.Millisecond)
	lastFrame := time.Now()
	for {
		// Update driver state first
		err = activeDriver.Update()
		if err != nil {
			break
		}

		err = game.Update()
		if err != nil {
			break
		}

		select {
		case <-drawChan:
			err = game.FixedUpdate(time.Now().Sub(lastFrame))
			if err != nil {
				break
			}
			lastFrame = time.Now()

			err = activeDriver.PreDraw()
			if err != nil {
				break
			}
			err = game.Draw()
			if err != nil {
				break
			}
			activeDriver.PostDraw()
		default:
			break
		}
		if err != nil {
			break
		}
	}
	return err
}
