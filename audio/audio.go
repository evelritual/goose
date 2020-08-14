package audio

// Player manages audio volume, channels, and overall playback.
type Player interface {
	NewSound(soundPath string) (Sound, error)
	SetVolume(volume float32) error
	Close() error
}

// Sound is a single sound bite.
type Sound interface {
	Play() error
	Close() error
}
