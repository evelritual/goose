package sdl2

import (
	"fmt"
	"io/ioutil"

	"github.com/veandco/go-sdl2/mix"

	"github.com/PapayaJuice/goose/audio"
)

const (
	maxVol = 128
	minVol = 0
)

// Player ...
type Player struct {
	currVol int
}

// Sound ...
type Sound struct {
	chunk *mix.Chunk
}

// NewAudioPlayer ...
func (s *SDL2) NewAudioPlayer() (audio.Player, error) {
	err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		return nil, fmt.Errorf("error opening sdl mix: %v", err)
	}

	return &Player{
		currVol: maxVol,
	}, nil
}

// SetVolume ...
func (p *Player) SetVolume(volume float32) error {
	if volume < 0.0 || volume > 1.0 {
		return fmt.Errorf("volume out of range")
	}

	v := int(float32(maxVol) * volume)
	// paranoia about floating point error
	if v < minVol {
		v = minVol
	} else if v > maxVol {
		v = maxVol
	}

	p.currVol = v
	mix.Volume(-1, v)
	return nil
}

// NewSound ...
func (p *Player) NewSound(soundPath string) (audio.Sound, error) {
	d, err := ioutil.ReadFile(soundPath)
	if err != nil {
		return nil, fmt.Errorf("error loading sound file: %v", err)
	}

	c, err := mix.QuickLoadWAV(d)
	if err != nil {
		return nil, fmt.Errorf("error loading wav: %v", err)
	}

	return &Sound{
		chunk: c,
	}, nil
}

// Close ...
func (p *Player) Close() error {
	mix.CloseAudio()
	return nil
}

// Play ...
func (s *Sound) Play() error {
	_, err := s.chunk.Play(-1, 0)
	if err != nil {
		return fmt.Errorf("error playing sound: %v", err)
	}
	return nil
}

// Close ...
func (s *Sound) Close() error {
	s.chunk.Free()
	return nil
}
