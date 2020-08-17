package main

import (
	"fmt"
	"log"
	"time"

	"github.com/PapayaJuice/goose"
	"github.com/PapayaJuice/goose/graphics"
)

var (
	playerFrame = 0
)

// Game ...
type Game struct {
	playerAtlus graphics.TextureAtlus
	frameTicker *time.Ticker
}

// Init ...
func (g *Game) Init() error {
	ta, err := goose.NewTextureAtlus("assets/player.png", 32, 32)
	if err != nil {
		return fmt.Errorf("error loading texture atlus: %v", err)
	}
	g.playerAtlus = ta
	g.frameTicker = time.NewTicker(1 * time.Second)

	goose.SetBackgroundColor(&graphics.ColorWhite)

	return nil
}

// Update ...
func (g *Game) Update() error {
	select {
	case <-g.frameTicker.C:
		playerFrame++
		if playerFrame >= g.playerAtlus.Len() {
			playerFrame = 0
		}
	default:
	}

	return nil
}

// Draw ...
func (g *Game) Draw() error {
	g.playerAtlus.Draw(
		playerFrame,
		(goose.GetWindowX()/2)-16,
		(goose.GetWindowY()/2)-16,
		1.0,
		1.0,
	)
	return nil
}

// Close ...
func (g *Game) Close() error {
	g.frameTicker.Stop()
	return g.playerAtlus.Close()
}

func main() {
	err := goose.Run(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
