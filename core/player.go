package core

import (
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
)

var img *ebiten.Image

func (p *player) Draw(screen *ebiten.Image) {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("data/assets/trainers_ow/frame1.png")
	if err != nil {
		log.Fatal(err)
	}
	screen.DrawImage(img, nil)
	ebitenutil.DebugPrintAt(screen, "player", p.pos.X, p.pos.Y)
}

type player struct {
	input *input.Handler
	pos   image.Point
}

func (p *player) Update() {
	if p.input.ActionIsPressed(ActionMoveLeft) {
		p.pos.X -= 4
		fmt.Printf("Moving to the left, current x is: %d\n", p.pos.X)
	}
	if p.input.ActionIsPressed(ActionMoveRight) {
		p.pos.X += 4
		fmt.Printf("Moving to the right, current x is: %d\n", p.pos.X)
	}
	if p.input.ActionIsPressed(ActionMoveDown) {
		p.pos.Y += 4
		fmt.Printf("Moving down, current y is: %d\n", p.pos.Y)
	}
	if p.input.ActionIsPressed(ActionMoveUp) {
		p.pos.Y -= 4
		fmt.Printf("Moving up, current y is: %d\n", p.pos.Y)
	}

}
