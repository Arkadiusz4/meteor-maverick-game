package main

import (
	"github.com/Arkadiusz4/meteor-maverick-game/game"
	_ "github.com/Arkadiusz4/meteor-maverick-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
	"os"
)

func main() {
	err := os.Setenv("DISPLAY", ":0")
	if err != nil {
		// Handle error if setting environment variable fails
		panic(err)
	}

	g := game.NewGame()

	err = ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
