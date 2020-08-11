package input

const (
	// Key0 "0"
	Key0 = Key(0)
	// Key1 "1"
	Key1 = Key(1)
	// Key2 "2"
	Key2 = Key(2)
	// Key3 "3"
	Key3 = Key(3)
	// Key4 "4"
	Key4 = Key(4)
	// Key5 "5"
	Key5 = Key(5)
	// Key6 "6"
	Key6 = Key(6)
	// Key7 "7"
	Key7 = Key(7)
	// Key8 "8"
	Key8 = Key(8)
	// Key9 "9"
	Key9 = Key(9)

	// KeyA "A"
	KeyA = Key(10)
	// KeyB "B"
	KeyB = Key(11)
	// KeyC "C"
	KeyC = Key(12)
	// KeyD "D"
	KeyD = Key(13)
	// KeyE "E"
	KeyE = Key(14)
	// KeyF "F"
	KeyF = Key(15)
	// KeyG "G"
	KeyG = Key(16)
	// KeyH "H"
	KeyH = Key(17)
	// KeyI "I"
	KeyI = Key(18)
	// KeyJ "J"
	KeyJ = Key(19)
	// KeyK "K"
	KeyK = Key(20)
	// KeyL "L"
	KeyL = Key(21)
	// KeyM "M"
	KeyM = Key(22)
	// KeyN "N"
	KeyN = Key(23)
	// KeyO "O"
	KeyO = Key(24)
	// KeyP "P"
	KeyP = Key(25)
	// KeyQ "Q"
	KeyQ = Key(26)
	// KeyR "R"
	KeyR = Key(27)
	// KeyS "S"
	KeyS = Key(28)
	// KeyT "T"
	KeyT = Key(29)
	// KeyU "U"
	KeyU = Key(30)
	// KeyV "V"
	KeyV = Key(31)
	// KeyW "W"
	KeyW = Key(32)
	// KeyX "X"
	KeyX = Key(33)
	// KeyY "Y"
	KeyY = Key(34)
	// KeyZ "Z"
	KeyZ = Key(35)

	// KeySpace "Space"
	KeySpace = Key(36)
	// KeyTab "Tab"
	KeyTab = Key(37)
	// KeyShiftL "Left Shift"
	KeyShiftL = Key(38)
	// KeyShiftR "Right Shift"
	KeyShiftR = Key(39)
	// KeyCtrlL "Left Ctrl/Control"
	KeyCtrlL = Key(40)
	// KeyCtrlR "Right Ctrl/Control"
	KeyCtrlR = Key(41)
	// KeyAltCmdL "Left Alt/Command"
	KeyAltCmdL = Key(42)
	// KeyAltCmdR "Left Alt/Command"
	KeyAltCmdR = Key(43)
	// KeyEnter "Enter"
	KeyEnter = Key(44)

	// KeyArrowLeft "Left Arrow"
	KeyArrowLeft = Key(45)
	// KeyArrowRight "Right Arrow"
	KeyArrowRight = Key(46)
	// KeyArrowDown "Down Arrow"
	KeyArrowDown = Key(47)
	// KeyArrowUp "Up Arrow"
	KeyArrowUp = Key(48)
)

// KeyState represents the current state of a key.
type KeyState struct {
	Pressed bool
	Repeat  bool
}
