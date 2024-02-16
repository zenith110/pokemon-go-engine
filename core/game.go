package core

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideHeight, outsideHeight
}

func CreateGame() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pokemon Go Engine")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
