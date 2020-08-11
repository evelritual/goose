package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/PapayaJuice/goose"
	"github.com/PapayaJuice/goose/graphics"
	"github.com/PapayaJuice/goose/input"
)

// Game ...
type Game struct {
	keyboard input.Keyboard

	bulletTexture graphics.Texture
	bullets       []*Bullet

	enemyTexture graphics.Texture
	enemyTicker  *time.Ticker
	enemies      []*Enemy

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
		x:     (goose.GetWindowX() / 2) - (pt.W() / 2),
		y:     goose.GetWindowY() - pt.H() - 5,
	}

	et, err := goose.NewTexture("assets/enemy.png")
	if err != nil {
		return fmt.Errorf("error loading enemy texture: %v", err)
	}
	g.enemyTexture = et
	g.enemyTicker = time.NewTicker(1 * time.Second)

	bt, err := goose.NewTexture("assets/bullet.png")
	if err != nil {
		return fmt.Errorf("error loading bullet texture: %v", err)
	}
	g.bulletTexture = bt

	return nil
}

// Close ...
func (g *Game) Close() error {
	g.enemyTicker.Stop()

	err := g.player.tex.Close()
	if err != nil {
		return fmt.Errorf("error closing player texture: %v", err)
	}

	err = g.enemyTexture.Close()
	if err != nil {
		return fmt.Errorf("error closing enemy texture: %v", err)
	}

	err = g.bulletTexture.Close()
	if err != nil {
		return fmt.Errorf("error closing bullet texture: %v", err)
	}

	return nil
}

// Update ...
func (g *Game) Update() error {
	// Input
	if g.keyboard.IsKeyDown(input.KeyArrowLeft) {
		g.player.x -= g.player.speed
		if g.player.x < 0 {
			g.player.x = 0
		}
	}
	if g.keyboard.IsKeyDown(input.KeyArrowRight) {
		g.player.x += g.player.speed
		if g.player.x > goose.GetWindowX()-g.player.tex.W() {
			g.player.x = goose.GetWindowX() - g.player.tex.W()
		}
	}

	if g.keyboard.IsKeyPress(input.KeySpace) {
		b := &Bullet{
			tex:   g.bulletTexture,
			speed: 5,
			x:     g.player.x + (g.bulletTexture.W() / 2),
			y:     g.player.y - g.bulletTexture.H(),
		}
		g.bullets = append(g.bullets, b)
	}

	// Enemies
	select {
	case <-g.enemyTicker.C:
		e := &Enemy{
			tex:   g.enemyTexture,
			speed: 2,
			x:     rand.Int31n(goose.GetWindowX()),
			y:     0,
		}
		g.enemies = append(g.enemies, e)
	default:
	}

	remEnemies := []int{}
	for i, e := range g.enemies {
		if e.destroyed || e.y > goose.GetWindowY() {
			remEnemies = append(remEnemies, i)
			continue
		}

		e.y += e.speed
	}

	for _, ei := range remEnemies {
		g.enemies[ei] = g.enemies[len(g.enemies)-1]
		g.enemies = g.enemies[:len(g.enemies)-1]
	}

	// Bullets
	remBullets := []int{}
	for i, b := range g.bullets {
		if b.destroyed || b.y < 0 {
			remBullets = append(remBullets, i)
			continue
		}

		b.y -= b.speed
	}

	for _, bi := range remBullets {
		g.bullets[bi] = g.bullets[len(g.bullets)-1]
		g.bullets = g.bullets[:len(g.bullets)-1]
	}

	// This is terrible collision detection but whatever
	for _, b := range g.bullets {
		for _, e := range g.enemies {
			if b.x > e.x+e.tex.W() {
				continue
			}
			if b.x+b.tex.W() < e.x {
				continue
			}
			if b.y > e.y+e.tex.H() {
				continue
			}
			if b.y+b.tex.H() < e.y {
				continue
			}
			// Otherwise, collision
			b.destroyed = true
			e.destroyed = true
		}
	}

	return nil
}

// Draw ...
func (g *Game) Draw() error {
	err := g.player.tex.Draw(g.player.x, g.player.y, 1.0, 1.0)
	if err != nil {
		return fmt.Errorf("error drawing player texture: %v", err)
	}

	for _, e := range g.enemies {
		err = e.tex.Draw(e.x, e.y, 1.0, 1.0)
		if err != nil {
			return fmt.Errorf("error drawing enemy texture: %v", err)
		}
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
