package main

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	_ "github.com/Arkadiusz4/meteor-maverick-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

type Game struct {
	player *game.Player
}

func (g *Game) Update() error {
	g.player.Update()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
}
