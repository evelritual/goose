package audio

// Player ...
type Player interface {
	NewSound(soundPath string) (Sound, error)
	SetVolume(volume float32) error
	Close() error
}

// Sound ...
type Sound interface {
	Play() error
	Close() error
}
