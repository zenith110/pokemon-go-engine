package core

import (
	"image"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
	input "github.com/quasilyte/ebitengine-input"
)

type GameControls struct {
	Controls []struct {
		Action string       `toml:"action"`
		Keys   []string     `toml:"keys"`
		Axis   string       `toml:"axis"`
		Value  input.Action `toml:"value"`
	} `toml:"controls"`
}

var controls GameControls

const (
	ActionMoveLeft input.Action = iota
	ActionMoveRight
	ActionMoveUp
	ActionMoveDown
)

func SetUpControls(game *Game) {
	file, err := os.Open("core/controls.toml")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = toml.Unmarshal(b, &controls)
	if err != nil {
		panic(err)
	}
	game.inputSystem.Init(input.SystemConfig{
		DevicesEnabled: input.AnyDevice,
	})

	keymap := input.Keymap{
		ActionMoveLeft:  {input.KeyGamepadLeft, input.KeyLeft, input.KeyA},
		ActionMoveRight: {input.KeyGamepadRight, input.KeyRight, input.KeyD},
		ActionMoveDown:  {input.KeyGamepadDown, input.KeyDown, input.KeyS},
		ActionMoveUp:    {input.KeyGamepadUp, input.KeyUp, input.KeyW},
	}
	game.p = &player{
		input: game.inputSystem.NewHandler(0, keymap),
		pos:   image.Point{X: 96, Y: 96},
	}
}
