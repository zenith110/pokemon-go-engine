package core

import input "github.com/quasilyte/ebitengine-input"

// actionMap maps action strings to input.Action constants
var actionMap = map[string]input.Action{
	"left":  ActionMoveLeft,
	"right": ActionMoveRight,
	"up":    ActionMoveUp,
	"down":  ActionMoveDown,
}

// keyMap maps string key names to input.Key constants
var keyMap = map[string]input.Key{
	"KeyGamePadLeft":  input.KeyGamepadLeft,
	"KeyGamePadRight": input.KeyGamepadRight,
	"KeyGamePadDown":  input.KeyGamepadDown,
	"KeyGamePadUp":    input.KeyGamepadUp,
	"KeyLeft":         input.KeyLeft,
	"KeyRight":        input.KeyRight,
	"KeyDown":         input.KeyDown,
	"KeyUp":           input.KeyUp,
	"KeyA":            input.KeyA,
	"KeyD":            input.KeyD,
	"KeyS":            input.KeyS,
	"KeyW":            input.KeyW,
}
