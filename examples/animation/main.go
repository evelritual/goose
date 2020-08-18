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
	playerAtlas graphics.TextureAtlas
	frameTicker *time.Ticker
}

// Init ...
func (g *Game) Init() error {
	ta, err := goose.NewTextureAtlas("assets/player.png", 32, 32)
	if err != nil {
		return fmt.Errorf("error loading texture atlas: %v", err)
	}
	g.playerAtlas = ta
	g.frameTicker = time.NewTicker(1 * time.Second)

	goose.SetBackgroundColor(&graphics.ColorWhite)

	return nil
}

// Update ...
func (g *Game) Update() error {
	select {
	case <-g.frameTicker.C:
		playerFrame++
		if playerFrame >= g.playerAtlas.Len() {
			playerFrame = 0
		}
	default:
	}

	return nil
}

// Draw ...
func (g *Game) Draw() error {
	g.playerAtlas.Draw(
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
	return g.playerAtlas.Close()
}

func main() {
	err := goose.Run(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
