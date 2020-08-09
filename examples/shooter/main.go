package main

import (
	"fmt"
	"log"

	"github.com/PapayaJuice/goose/internal/texture"

	"github.com/PapayaJuice/goose/internal/input"

	"github.com/PapayaJuice/goose"
)

// Game ...
type Game struct {
	keyboard input.Keyboard

	bulletTexture texture.Texture
	bullets       []*Bullet

	player *Player
}

// Init ...
func (g *Game) Init() error {
	g.keyboard = goose.NewKeyboard()

	pt, err := goose.NewTexture("assets/player.png")
	if err != nil {
		return fmt.Errorf("error loading player texture: %v", err)
	}

	g.player = &Player{
		tex:   pt,
		speed: 3,
	}

	bt, err := goose.NewTexture("assets/bullet.png")
	if err != nil {
		return fmt.Errorf("error loading bullet texture: %v", err)
	}
	g.bulletTexture = bt

	return nil
}

// Close ...
func (g *Game) Close() error {
	err := g.player.tex.Close()
	if err != nil {
		return fmt.Errorf("error closing player texture: %v", err)
	}

	err = g.bulletTexture.Close()
	if err != nil {
		return fmt.Errorf("error closing bullet texture: %v", err)
	}
	return nil
}

// Update ...
func (g *Game) Update() error {
	if g.keyboard.IsKeyDown(input.KeyArrowLeft) {
		g.player.x -= g.player.speed
	}
	if g.keyboard.IsKeyDown(input.KeyArrowRight) {
		g.player.x += g.player.speed
	}
	if g.keyboard.IsKeyDown(input.KeyArrowUp) {
		g.player.y -= g.player.speed
	}
	if g.keyboard.IsKeyDown(input.KeyArrowDown) {
		g.player.y += g.player.speed
	}

	if g.keyboard.IsKeyPress(input.KeySpace) {
		b := &Bullet{
			tex:   g.bulletTexture,
			speed: 5,
			x:     g.player.x,
			y:     g.player.y,
		}
		g.bullets = append(g.bullets, b)
	}

	remBullets := []int{}
	for i, b := range g.bullets {
		if b.destroyed {
			remBullets = append(remBullets, i)
			continue
		}

		b.y -= b.speed
	}

	for _, bi := range remBullets {
		g.bullets[bi] = g.bullets[len(g.bullets)-1]
		g.bullets = g.bullets[:len(g.bullets)-1]
	}

	return nil
}

// Draw ...
func (g *Game) Draw() error {
	err := g.player.tex.Draw(g.player.x, g.player.y, 1.0, 1.0)
	if err != nil {
		return fmt.Errorf("error drawing player texture: %v", err)
	}

	for _, b := range g.bullets {
		err = b.tex.Draw(b.x, b.y, 1.0, 1.0)
		if err != nil {
			return fmt.Errorf("error drawing bullet texture: %v", err)
		}
	}
	return nil
}

func main() {
	err := goose.Run(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
