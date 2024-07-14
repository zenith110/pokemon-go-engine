package core

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	input "github.com/quasilyte/ebitengine-input"
)

type DiscordRichPresence struct {
}

var drp DiscordRichPresence

type Game struct {
	p           *player
	inputSystem input.System
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// g.p.Draw(screen)
	ebitenutil.DebugPrintAt(screen, "New Game", 160, 120)
	ebitenutil.DebugPrintAt(screen, "Options", 160, 140)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideHeight, outsideHeight
}

func CreateGame(title string, width int, height int) {

	game := &Game{}
	drp.SetupRichPresence()
	data := map[string]string{"state": "Idle", "details": "In home screen", "largeImage": "", "largeText": "", "inBattle": "false"}
	drp.UpdateRichPresence(data)
	// SetUpControls(game)
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Pokemon Go Engine")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
