package goose

import "github.com/PapayaJuice/goose/audio"

// NewAudioPlayer ...
func NewAudioPlayer() (audio.Player, error) {
	return activeDriver.NewAudioPlayer()
}
