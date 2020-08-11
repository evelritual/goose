package sdl2

import (
	"github.com/veandco/go-sdl2/sdl"

	"github.com/PapayaJuice/goose/input"
)

var (
	keyMap = map[sdl.Keycode]input.Key{
		sdl.K_0:      input.Key0,
		sdl.K_1:      input.Key1,
		sdl.K_2:      input.Key2,
		sdl.K_3:      input.Key3,
		sdl.K_4:      input.Key4,
		sdl.K_5:      input.Key5,
		sdl.K_6:      input.Key6,
		sdl.K_7:      input.Key7,
		sdl.K_8:      input.Key8,
		sdl.K_9:      input.Key9,
		sdl.K_a:      input.KeyA,
		sdl.K_b:      input.KeyB,
		sdl.K_c:      input.KeyC,
		sdl.K_d:      input.KeyD,
		sdl.K_e:      input.KeyE,
		sdl.K_f:      input.KeyF,
		sdl.K_g:      input.KeyG,
		sdl.K_h:      input.KeyH,
		sdl.K_i:      input.KeyI,
		sdl.K_j:      input.KeyJ,
		sdl.K_k:      input.KeyK,
		sdl.K_l:      input.KeyL,
		sdl.K_m:      input.KeyM,
		sdl.K_n:      input.KeyN,
		sdl.K_o:      input.KeyO,
		sdl.K_p:      input.KeyP,
		sdl.K_q:      input.KeyQ,
		sdl.K_r:      input.KeyR,
		sdl.K_s:      input.KeyS,
		sdl.K_t:      input.KeyT,
		sdl.K_u:      input.KeyU,
		sdl.K_v:      input.KeyV,
		sdl.K_w:      input.KeyW,
		sdl.K_x:      input.KeyX,
		sdl.K_y:      input.KeyY,
		sdl.K_z:      input.KeyZ,
		sdl.K_SPACE:  input.KeySpace,
		sdl.K_TAB:    input.KeyTab,
		sdl.K_LSHIFT: input.KeyShiftL,
		sdl.K_RSHIFT: input.KeyShiftR,
		sdl.K_LCTRL:  input.KeyCtrlL,
		sdl.K_RCTRL:  input.KeyCtrlR,
		sdl.K_LALT:   input.KeyAltCmdL,
		sdl.K_RALT:   input.KeyAltCmdR,
		sdl.K_RETURN: input.KeyEnter,
		sdl.K_LEFT:   input.KeyArrowLeft,
		sdl.K_RIGHT:  input.KeyArrowRight,
		sdl.K_DOWN:   input.KeyArrowDown,
		sdl.K_UP:     input.KeyArrowUp,
	}
)
