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

// Player wraps needed methods for the audio.Player interface.
type Player struct {
	currVol int
}

// Sound holds SDL chunk data for playback in use with the audio.Sound interface.
type Sound struct {
	chunk *mix.Chunk
}

// NewAudioPlayer initializes the SDL mixer and returns a default Player.
// Player must be closed manually.
func (s *SDL2) NewAudioPlayer() (audio.Player, error) {
	err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		return nil, fmt.Errorf("error opening sdl mix: %v", err)
	}

	return &Player{
		currVol: maxVol,
	}, nil
}

// SetVolume sets the volume across the entire SDL player. This will not hold
// true if new mixing channels are allocated beyond the default.
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

// NewSound loads a WAV file into memory and converts it into a chunk. Sound
// must be closed manually.
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

// Close releases audio player resources and stops all playback.
func (p *Player) Close() error {
	mix.CloseAudio()
	return nil
}

// Play begins playback of the Sound through the sdl mixer.
func (s *Sound) Play() error {
	_, err := s.chunk.Play(-1, 0)
	if err != nil {
		return fmt.Errorf("error playing sound: %v", err)
	}
	return nil
}

// Close frees the sdl chunk resource.
func (s *Sound) Close() error {
	s.chunk.Free()
	return nil
}
