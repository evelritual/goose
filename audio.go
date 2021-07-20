package goose

import "github.com/evelritual/goose/audio"

// NewAudioPlayer initiliazes a new audio player for use in allocating sounds
// and dealing with playback.
func NewAudioPlayer() (audio.Player, error) {
	return activeDriver.NewAudioPlayer()
}
