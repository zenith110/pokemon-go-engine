package core

import (
	"image"
	"io"
	"os"

	"github.com/pelletier/go-toml/v2"
	input "github.com/quasilyte/ebitengine-input"
)

type Control struct {
	Action   string       `toml:"action"`
	InputKey string       `toml:"inputKey"`
	Keys     []string     `toml:"keys"`
	Axis     string       `toml:"axis"`
	Value    input.Action `toml:"value"`
}

type GameControls struct {
	Controls []Control `toml:"controls"`
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

	// Create dynamic keymap from TOML data
	keymap := input.Keymap{}

	for _, control := range controls.Controls {
		action, exists := actionMap[control.Action]
		if !exists {
			continue // Skip unknown actions
		}

		var keys []input.Key
		for _, keyStr := range control.Keys {
			if key, exists := keyMap[keyStr]; exists {
				keys = append(keys, key)
			}
		}

		if len(keys) > 0 {
			keymap[action] = keys
		}
	}

	game.p = &player{
		input: game.inputSystem.NewHandler(0, keymap),
		pos:   image.Point{X: 96, Y: 96},
	}
}
