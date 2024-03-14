package main

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"os"
)

func main() {
	err := os.Setenv("EBITEN_HEADLESS", "1")
	if err != nil {
		panic(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Headless Mode")

	g := game.NewGame()

	err = ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
