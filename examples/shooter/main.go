package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/PapayaJuice/goose"
	"github.com/PapayaJuice/goose/audio"
	"github.com/PapayaJuice/goose/graphics"
	"github.com/PapayaJuice/goose/input"
)

// Game ...
type Game struct {
	audioPlayer audio.Player
	keyboard    input.Keyboard
	font        graphics.Font
	score       int

	bulletTexture graphics.Texture
	bullets       []*Bullet

	enemyTexture      graphics.Texture
	enemyDestroySound audio.Sound
	enemyTicker       *time.Ticker
	enemies           []*Enemy

	player *Player
}

// Init ...
func (g *Game) Init() error {
	ap, err := goose.NewAudioPlayer()
	if err != nil {
		return fmt.Errorf("error initializing audio player: %v", err)
	}
	err = ap.SetVolume(0.2)
	if err != nil {
		return fmt.Errorf("error setting volume: %v", err)
	}
	g.audioPlayer = ap

	s, err := ap.NewSound("assets/hit.wav")
	if err != nil {
		return fmt.Errorf("error loading enemy sound: %v", err)
	}
	g.enemyDestroySound = s

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

	f, err := goose.NewFont("assets/font.ttf", 16)
	if err != nil {
		return fmt.Errorf("error loading font: %v", err)
	}
	g.font = f

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

	err = g.font.Close()
	if err != nil {
		return fmt.Errorf("error closing font: %v", err)
	}

	err = g.audioPlayer.Close()
	if err != nil {
		return fmt.Errorf("error closing audio player: %v", err)
	}
	err = g.enemyDestroySound.Close()
	if err != nil {
		return fmt.Errorf("error closing enemy sound: %v", err)
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
			x:     rand.Int31n(goose.GetWindowX() - g.enemyTexture.W()),
			y:     0,
		}
		g.enemies = append(g.enemies, e)
	default:
	}

	newEnemies := []*Enemy{}
	for _, e := range g.enemies {
		if e.destroyed || e.y > goose.GetWindowY() {
			continue
		}

		newEnemies = append(newEnemies, e)
		e.y += e.speed
	}
	g.enemies = newEnemies

	// Bullets
	newBullets := []*Bullet{}
	for _, b := range g.bullets {
		if b.destroyed || b.y < 0 {
			continue
		}

		newBullets = append(newBullets, b)
		b.y -= b.speed
	}
	g.bullets = newBullets

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
			g.score += 10

			err := g.enemyDestroySound.Play()
			if err != nil {
				return fmt.Errorf("error playing sound: %v", err)
			}
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

	// UI
	err = g.font.Draw(fmt.Sprintf("Score: %d", g.score), 5, 0, graphics.ColorWhite)
	if err != nil {
		return fmt.Errorf("error drawing font: %v", err)
	}

	return nil
}

func main() {
	err := goose.Run(&Game{})
	if err != nil {
		log.Fatal(err)
	}
}
